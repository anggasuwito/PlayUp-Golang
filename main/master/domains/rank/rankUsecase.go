package rank

import (
	"database/sql"
	"playUp/main/master/model"
	"sync"
)

type RankUsecase struct {
	db       *sql.DB
	RankRepo RankRepository
}

type RankUsecaseInterface interface {
	getRankByIDSport(res *model.Result, wg *sync.WaitGroup, id string)
}

func NewRankUsecase(db *sql.DB) RankUsecaseInterface {
	return RankUsecase{db, NewRankRepo(db)}
}
func (r RankUsecase) getRankByIDSport(res *model.Result, wg *sync.WaitGroup, id string) {
	defer wg.Done()

	err := r.RankRepo.getRankByIDSport(res, id)
	if err != nil {
		res.Err = err
	}
}
