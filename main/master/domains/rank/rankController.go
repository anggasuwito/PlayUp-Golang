package rank

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"playUp/main/master/model"
	"playUp/main/master/utils/helper"
	"sync"
)

type Controller struct {
	RankUsecase RankUsecaseInterface
}

func NewController(db *sql.DB) *Controller {
	return &Controller{RankUsecase: NewRankUsecase(db)}
}

func (c Controller) HandleGetRankBySportID(w http.ResponseWriter, r *http.Request) {

	var result model.Result
	var wg sync.WaitGroup
	id := helper.DecodeParams("id", r)

	wg.Add(1)
	go c.RankUsecase.getRankByIDSport(&result,&wg, id)
	wg.Wait()
	if result.Err != nil {
		result := model.Response{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: "Ambil rank gagal",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)

		w.Write([]byte(byteOfResult))
	} else {
		result := model.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Ambil rank berhasil",
			Data:    result.Res,
		}
		byteOfResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteOfResult))
	}

}
