package badmintonSingle

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"playUp/main/master/domains/user"
	"playUp/main/master/model"

	"github.com/gorilla/mux"
)

type Controller struct {
	BadmintonSUsecase BadmintonSUsecaseInterface
}

func NewController(db *sql.DB) *Controller {
	return &Controller{BadmintonSUsecase: NewBadmintonSUsecase(db)}
}


func (c Controller) HandleFindOpponent(w http.ResponseWriter, r *http.Request) {

	var player user.User

	_ = json.NewDecoder(r.Body).Decode(&player)

	if len(handlehit) == 0 {
		fmt.Println("ini 0")
		handlehit = append(handlehit, player)
	} else if len(handlehit) == 1 {
		fmt.Println("ini 1")
		for _, value := range handlehit {
			if value.ID != player.ID {
				handlehit = append(handlehit, player)
			}
		}
	}

	err := c.BadmintonSUsecase.createRoomBadmintonS(&PlayerList, &player, &MatchID)

	if err != nil {
		result := model.Response{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: "lawan tidak ditemukan ",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)

		w.Write([]byte(byteOfResult))
	} else if len(PlayerList) == 1 {
		result := model.Response{
			Status:  "pending",
			Code:    http.StatusNotFound,
			Message: "Lawan belum ditemukan",
			Data: map[string]interface{}{
				"match_id":      MatchID,
				"match_players": PlayerList,
			},
		}
		byteOfResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteOfResult))
	} else {
		result := model.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Lawan ditemukan",
			Data: map[string]interface{}{
				"match_id":      MatchID,
				"match_players": PlayerList,
			},
		}
		byteOfResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteOfResult))

	}

}
func (c Controller) HandleCancelMatch(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	idUser := param["id"]
	for i, v := range PlayerList {
		if idUser == v.ID {
			PlayerList = append(PlayerList[:i], PlayerList[i+1:]...)
			handlehit = append(handlehit[:i], handlehit[i+1:]...)
		}
	}
}
func (c Controller) HandleResetRoom(w http.ResponseWriter, r *http.Request) {
	if len(handlehit) < 2 {
		result := model.Response{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: "gagal reset, kurang pemain ",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)

		w.Write([]byte(byteOfResult))
	} else {
		result := model.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "berhasil reset ",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)

		w.Write([]byte(byteOfResult))
		PlayerList = []user.User{}
		handlehit = []user.User{}
		MatchID = ""
	}
}