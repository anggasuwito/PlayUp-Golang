package badmintonSingle

import (
	"database/sql"
	"log"
	"playUp/main/master/domains/user"
)

type BadmintonSUsecase struct {
	db     *sql.DB
	BSRepo BadmintonSRepository
}
type BadmintonSUsecaseInterface interface {
	createRoomBadmintonS(playersList *[]user.User, players *user.User, matchID *string) error
}

func NewBadmintonSUsecase(db *sql.DB) BadmintonSUsecase {
	return BadmintonSUsecase{db, NewBadmintonSRepo(db)}
}
func (r BadmintonSUsecase) createRoomBadmintonS(playersList *[]user.User, player *user.User, matchID *string) error {

	if len(*playersList) == 0 {
		*playersList = append(*playersList, *player)
		return nil
	} else if len(*playersList) == 1 {
		for _, value := range PlayerList {
			if value.ID == player.ID {
				return nil
			}
		}
		*playersList = append(*playersList, *player)
	} else if len(*playersList) == 2 {
		return nil
	}
	r.BSRepo.createMatch(*playersList, matchID)
	log.Println("generate match id")
	return nil

}
