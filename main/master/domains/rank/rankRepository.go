package rank

import (
	"database/sql"
	"log"
	"playUp/main/master/model"
	query "playUp/main/master/utils/db"
)

type RankRepo struct {
	db *sql.DB
}

type RankRepository interface {
	getRankByIDSport(res *model.Result, id string) error
}

func NewRankRepo(db *sql.DB) RankRepository {
	return RankRepo{db: db}
}

func (d RankRepo) getRankByIDSport(res *model.Result, id string) error {
	var rank []*rankList
	tx, err := d.db.Begin()

	if err != nil {
		log.Printf("ERROR Failed get rank by sport id 1 %v \n", err)
		return err
	}
	rows, err := tx.Query(query.SELECT_RANK_BY_ID_SPORT, id)

	if err != nil {
		log.Printf("ERROR Failed get rank by sport id 2 %v \n", err)
		tx.Rollback()
		return err
	}
	for rows.Next() {
		var regRank rankList
		err := rows.Scan(&regRank.RankUserID, &regRank.RankUserPhotoProfile, &regRank.RankUserUserName, &regRank.RankUserUserFullName, &regRank.RankUserRankID, &regRank.RankUserMatchCount, &regRank.RankUserGradeCount, &regRank.RankUserMatchSportID)
		if err != nil {
			log.Printf("ERROR Failed get rank by sport id 3 %v \n", err)
			tx.Rollback()
			return err
		}
		rank = append(rank, &regRank)
	}

	res.Res = rank
	tx.Commit()
	return nil
}
