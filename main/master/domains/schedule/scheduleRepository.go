package schedule

import (
	"database/sql"
	"errors"
	"log"
	"playUp/main/master/domains/user"
	"playUp/main/master/model"
	query "playUp/main/master/utils/db"
	generate "playUp/main/master/utils/generator"
)

type ScheduleRepo struct {
	db *sql.DB
}

type ScheduleRepository interface {
	createScheduleBadmintonS(ns *schedule, res *model.Result) error

	updateScheduleBadmintonS(newSchedule *schedule, finalResult *resultMatch) error

	getScheduleBadmintonS(scheduleList *[]*schedule, id, status string) error

	getScheduleBadmintonSALL(res *model.Result, page, limit *string) error

	getScheduleDetailBadmintonS(res *model.Result, id *string) error
}

func NewScheduleRepo(db *sql.DB) ScheduleRepository {
	return ScheduleRepo{db: db}
}

// createScheduleBadmintonS -> create new badminton schedule
func (d ScheduleRepo) createScheduleBadmintonS(ns *schedule, res *model.Result) error {
	regSchedule := new(schedule)
	scheduleID := generate.GenerateUUID()
	bsplID := generate.GenerateUUID()
	tx, err := d.db.Begin()

	if err != nil {
		return err
	}

	err = tx.QueryRow(query.SELECT_SCHEDULE_BY_MATCH_ID, ns.ScheduleMatchID).Scan(&regSchedule.ScheduleID, &regSchedule.ScheduleMatchID, &regSchedule.ScheduleUserID, &regSchedule.ScheduleUserName, &regSchedule.SchedeleLocation, &regSchedule.ScheduleTime, &regSchedule.ScheduleStatus, &regSchedule.ScheduleResult, &regSchedule.ScheduleOpponent, &regSchedule.ScheduleOpponentID)

	if err == nil {
		log.Printf("-> error failed create badminton schedule 1 %v \n", err)
		tx.Rollback()
		return errors.New("sudah ada datanya")
	}

	stmt, _ := tx.Prepare(query.INSERT_SCHEDULE)

	_, err = stmt.Exec(scheduleID, ns.ScheduleMatchID, ns.ScheduleUserID, ns.ScheduleUserName, ns.SchedeleLocation, ns.ScheduleTime, ns.ScheduleStatus, ns.ScheduleResult, ns.ScheduleOpponent, ns.ScheduleOpponentID)

	if err != nil {
		log.Printf("-> error failed create badminton schedule 2 %v \n", err)
		tx.Rollback()
		return err
	}
	stmt2, _ := tx.Prepare(query.INSERT_BADMINTON_SINGLE_PLAYER_LIST)

	_, err = stmt2.Exec(bsplID, ns.ScheduleMatchID, ns.ScheduleUserName, ns.ScheduleOpponent)

	if err != nil {
		log.Printf("-> error failed create badminton schedule 3 %v \n", err)
		tx.Rollback()
		return err
	}
	err = tx.QueryRow(query.SELECT_SCHEDULE_BY_ID, scheduleID).Scan(&ns.ScheduleID, &ns.ScheduleMatchID, &ns.ScheduleUserID, &ns.ScheduleUserName, &ns.SchedeleLocation, &ns.ScheduleTime, &ns.ScheduleStatus, &ns.ScheduleResult, &ns.ScheduleOpponent, &ns.ScheduleOpponentID)
	if err != nil {
		log.Printf("-> error failed create badminton schedule 4 %v \n", err)
		tx.Rollback()
		return err

	}
	ns.ScheduleID = scheduleID
	res.Res = ns
	tx.Commit()
	return nil
}

func (d ScheduleRepo) updateScheduleBadmintonS(nSchedule *schedule, finalResult *resultMatch) error {
	regSchedule := new(schedule)

	tx, err := d.db.Begin()

	if err != nil {
		log.Println(err)
		return err
	}
	err = tx.QueryRow(query.SELECT_SCHEDULE_BY_MATCH_ID_AND_STATUS, nSchedule.ScheduleMatchID, "I").Scan(&regSchedule.ScheduleID, &regSchedule.ScheduleMatchID, &regSchedule.ScheduleUserID, &regSchedule.ScheduleUserName, &regSchedule.SchedeleLocation, &regSchedule.ScheduleTime, &regSchedule.ScheduleStatus, &regSchedule.ScheduleResult, &regSchedule.ScheduleOpponent, &regSchedule.ScheduleOpponentID)

	log.Println(regSchedule.ScheduleStatus)

	if err == nil || regSchedule.ScheduleStatus == "I" {
		log.Printf("schedule sudah diupdate -> %v \n", err)
		tx.Rollback()
		return errors.New("schedule sudah diupdate")
	}

	stmt, _ := tx.Prepare(query.UPDATE_SCHEDULE)
	_, err = stmt.Exec(nSchedule.ScheduleResult, nSchedule.ScheduleID)

	if err != nil {
		log.Printf("-> error failed update schedule 1 %v \n", err)
		tx.Rollback()
		return err
	}

	stmt2, _ := tx.Prepare(query.UPDATE_RESULT)

	_, err = stmt2.Exec(finalResult.Winner, finalResult.Loser, nSchedule.ScheduleMatchID)
	if err != nil {
		log.Printf("-> error failed update schedule 2 %v \n", err)
		tx.Rollback()
		return err
	}
	regRank := new(rank)

	_ = tx.QueryRow(query.SELECT_RANK_BY_USER_ID, finalResult.WinnerID).Scan(&regRank.rankMatchCount, &regRank.rankGradeCount)

	matchCount := regRank.rankMatchCount + 1
	gradeCount := regRank.rankGradeCount + 1

	stmt3, err := tx.Prepare(query.UPDATE_RANK)

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = stmt3.Exec(matchCount, gradeCount, finalResult.WinnerID)

	if err != nil {
		log.Printf("-> error failed update schedule 3 %v \n", err)
		tx.Rollback()
		return errors.New("gagal update rank winner")
	}

	rank := new(rank)

	_ = tx.QueryRow(query.SELECT_RANK_BY_USER_ID, finalResult.LoserID).Scan(&rank.rankMatchCount, &rank.rankGradeCount)

	rankMatchCount := rank.rankMatchCount + 1
	rankGradeCount := rank.rankGradeCount

	stmt4, _ := tx.Prepare(query.UPDATE_RANK)

	_, err = stmt4.Exec(rankMatchCount, rankGradeCount, finalResult.LoserID)

	if err != nil {
		log.Printf("-> error failed update schedule 4 %v \n", err)
		tx.Rollback()
		return errors.New("gagal update rank loser")
	}

	tx.Commit()
	return nil
}

// getScheduleBadmintonS -> get schedule badminton per user id
func (d ScheduleRepo) getScheduleBadmintonS(scheduleList *[]*schedule, id, status string) error {
	tx, err := d.db.Begin()

	if err != nil {
		log.Printf("-> error failed get schedule 1 %v \n", err)
		return err
	}
	rows, err := tx.Query(query.SELECT_SCHEDULE_BY_USER_OPPONENT_AND_STATUS, id, id, status)
	if err != nil {
		log.Printf("-> error failed get schedule 2 %v \n", err)
		tx.Rollback()
		return err
	}
	for rows.Next() {
		var regSchedule schedule
		err := rows.Scan(&regSchedule.ScheduleID, &regSchedule.ScheduleMatchID, &regSchedule.ScheduleUserID, &regSchedule.ScheduleUserName, &regSchedule.SchedeleLocation, &regSchedule.ScheduleTime, &regSchedule.ScheduleStatus, &regSchedule.ScheduleResult, &regSchedule.ScheduleOpponent, &regSchedule.ScheduleOpponentID)
		if err != nil {
			return errors.New("FAILED SCAN")
		}
		*scheduleList = append(*scheduleList, &regSchedule)
	}
	tx.Commit()
	return nil
}

// getScheduleDetailBadmintonS -> get schedule detail badminton by schedule id
func (d ScheduleRepo) getScheduleDetailBadmintonS(res *model.Result, id *string) error {

	detailSchedule := new(schedule)
	user1 := new(user.User)
	user2 := new(user.User)

	var playerOnePhoto string
	var playerTwoPhoto string

	err := d.db.QueryRow(query.SELECT_SCHEDULE_BY_ID, *id).Scan(&detailSchedule.ScheduleID, &detailSchedule.ScheduleMatchID, &detailSchedule.ScheduleUserID, &detailSchedule.ScheduleUserName, &detailSchedule.SchedeleLocation, &detailSchedule.ScheduleTime, &detailSchedule.ScheduleStatus, &detailSchedule.ScheduleResult, &detailSchedule.ScheduleOpponent, &detailSchedule.ScheduleOpponentID)

	if err != nil {
		log.Printf("-> error failed get detail schedule 1 %v \n", err)
		return err
	}
	err = d.db.QueryRow(query.SELECT_USER_BY_ID, detailSchedule.ScheduleUserID).Scan(&user1.ID, &user1.PhotoProfile, &user1.Username, &user1.UserFullName, &user1.UserGender, &user1.UserEmail, &user1.Password, &user1.Created, &user1.Updated, &user1.FacebookAccount, &user1.GoogleAccount)

	if err != nil {
		log.Printf("-> error failed get detail schedule 2 %v \n", err)
		return err
	}

	err = d.db.QueryRow(query.SELECT_USER_BY_ID, detailSchedule.ScheduleOpponentID).Scan(&user2.ID, &user2.PhotoProfile, &user2.Username, &user2.UserFullName, &user2.UserGender, &user2.UserEmail, &user2.Password, &user2.Created, &user2.Updated, &user2.FacebookAccount, &user2.GoogleAccount)

	if err != nil {
		log.Printf("-> error failed get detail schedule 3 %v \n", err)
		return err
	}

	playerOnePhoto = user1.PhotoProfile
	playerTwoPhoto = user2.PhotoProfile

	var resultSchedule schedule = schedule{
		ScheduleID:            detailSchedule.ScheduleID,
		ScheduleMatchID:       detailSchedule.ScheduleMatchID,
		ScheduleUserID:        detailSchedule.ScheduleUserID,
		ScheduleUserName:      detailSchedule.ScheduleUserName,
		SceduleUserPhoto:      playerOnePhoto,
		SchedeleLocation:      detailSchedule.SchedeleLocation,
		ScheduleTime:          detailSchedule.ScheduleTime,
		ScheduleStatus:        detailSchedule.ScheduleStatus,
		ScheduleResult:        detailSchedule.ScheduleResult,
		ScheduleOpponent:      detailSchedule.ScheduleOpponent,
		ScheduleOpponentID:    detailSchedule.ScheduleOpponentID,
		ScheduleOpponentPhoto: playerTwoPhoto,
	}

	res.Res = resultSchedule

	return nil
}

// getScheduleBadmintonSALL -> get all badminton schedule
func (d ScheduleRepo) getScheduleBadmintonSALL(res *model.Result, page, limit *string) error {
	var scheduleList []*schedule
	var totalData model.Total
	err := d.db.QueryRow(query.SELECT_SCHEDULE_TOTAL_DATA).Scan(&totalData)
	if err != nil {
		log.Printf("error failed get schedule badminton 1 %v \n", err)
		return err
	}
	rows, err := d.db.Query(query.SELECT_SCHEDULE_ALL, *page, *limit)
	if err != nil {
		log.Printf("error failed get schedule badminton 2 %v \n", err)
		return err
	}
	for rows.Next() {
		var regSchedule schedule
		err := rows.Scan(&regSchedule.ScheduleID, &regSchedule.ScheduleMatchID, &regSchedule.ScheduleUserID, &regSchedule.ScheduleUserName, &regSchedule.SchedeleLocation, &regSchedule.ScheduleTime, &regSchedule.ScheduleStatus, &regSchedule.ScheduleResult, &regSchedule.ScheduleOpponent, &regSchedule.ScheduleOpponentID)
		if err != nil {
			log.Println(err)
			return err

		}
		scheduleList = append(scheduleList, &regSchedule)
	}

	res.Res = map[string]interface{}{
		"total_data": totalData,
		"result":     scheduleList,
	}
	return nil
}
