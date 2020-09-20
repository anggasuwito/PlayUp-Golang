package badmintonSingle

import (
	"database/sql"
	"log"
	"playUp/main/master/domains/user"
	query "playUp/main/master/utils/db"
	generate "playUp/main/master/utils/generator"
)

type BadmintonSRepo struct {
	db *sql.DB
}
type BadmintonSRepository interface {
	createMatch(playerList []user.User, matchIDPrep *string) error
}

func NewBadmintonSRepo(db *sql.DB) BadmintonSRepository {
	return BadmintonSRepo{db: db}
}

func (d BadmintonSRepo) createMatch(playerList []user.User, matchIDPrep *string) error {

	matchID := generate.GenerateUUID()

	playerListMatchID := generate.GenerateUUID()

	tx, err := d.db.Begin()

	if err != nil {
		log.Println(err)
		return err
	}

	stmt, _ := tx.Prepare(query.INSERT_MATCH)

	_, err = stmt.Exec(matchID, 1)
	if err != nil {
		log.Println(err)
		tx.Rollback()
	}
	stmt2, _ := tx.Prepare(query.INSERT_BADMINTON_SINGLE_PLAYER_LIST)

	_, err = stmt2.Exec(playerListMatchID, matchID, playerList[0].Username, playerList[1].Username)
	if err != nil {
		log.Println("ERROR 2")
		log.Println(err)
		tx.Rollback()
	}

	stmt3, _ := tx.Prepare(query.INSERT_RESULT)

	_, err = stmt3.Exec(generate.GenerateUUID(), matchID, "unknown", "unknown")
	if err != nil {
		log.Println("ERROR 2")
		log.Println(err)
		tx.Rollback()
	}
	if *matchIDPrep == "" {
		*matchIDPrep = matchID
		log.Println("match id is empty")
		log.Print(*matchIDPrep)
	} else {
		*matchIDPrep = *matchIDPrep
		log.Println("match id is not empty")
	}
	tx.Commit()

	return nil

}
