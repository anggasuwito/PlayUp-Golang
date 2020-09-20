package chat
import (
	"database/sql"
	"encoding/json"
	"net/http"
	"playUp/main/master/model"
	"playUp/main/master/utils/helper"
)

type Controller struct {
	ChatUsecase ChatUsecaseInterface
}

func NewController(db *sql.DB) *Controller {
	return &Controller{ChatUsecase: NewChatUsecase(db)}
}

func (c Controller) HandleSendChat(w http.ResponseWriter, r *http.Request) {

	var newChat chat

	_ = json.NewDecoder(r.Body).Decode(&newChat)

	err := c.ChatUsecase.handleSendChat(&newChat)
	if err != nil {
		result := model.Response{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: "kirim Chat gagal",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)

		w.Write([]byte(byteOfResult))
	} else {
		result := model.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "kirim Chat berhasil",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteOfResult))
	}

}
func (c Controller) HandleGetChat(w http.ResponseWriter, r *http.Request) {

	var chatList []*chat
	
	MatchID := helper.DecodeParams("id",r)

	err := c.ChatUsecase.handleGetChat(&chatList,MatchID)
	if err != nil {
		result := model.Response{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: "Ambil Chat gagal",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)

		w.Write([]byte(byteOfResult))
	} else {
		result := model.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Ambil Chat berhasil",
			Data:    chatList,
		}
		byteOfResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteOfResult))
	}

}
