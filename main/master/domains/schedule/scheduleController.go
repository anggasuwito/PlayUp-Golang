package schedule

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"playUp/main/master/model"
	"playUp/main/master/utils/helper"
	"strconv"
	"sync"
)

type Controller struct {
	ScheduleUsecase ScheduleUsecaseInterface
}

func NewController(db *sql.DB) *Controller {
	return &Controller{ScheduleUsecase: NewScheduleUsecase(db)}
}

func (c Controller) HandleCreateSchedule(w http.ResponseWriter, r *http.Request) {

	var wg sync.WaitGroup
	var newSchedule schedule
	var result model.Result

	_ = json.NewDecoder(r.Body).Decode(&newSchedule)

	wg.Add(1)
	go c.ScheduleUsecase.createScheduleBadmintonS(&newSchedule, &result, &wg)
	wg.Wait()
	if result.Err != nil {
		result := model.Response{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: "Buat jadwal baru gagal",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)

		w.Write([]byte(byteOfResult))
	} else {
		result := model.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Buat jadwal baru berhasil",
			Data:    result.Res,
		}
		byteOfResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteOfResult))
	}

}
func (c Controller) HandleUpdateSchedule(w http.ResponseWriter, r *http.Request) {

	var newSchedule schedule
	var result model.Result
	var wg sync.WaitGroup

	_ = json.NewDecoder(r.Body).Decode(&newSchedule)

	var finalResult = resultMatch{
		Winner:   "",
		WinnerID: "",
		Loser:    "",
		LoserID:  "",
	}

	if newSchedule.ScheduleResult == "W" {
		finalResult.Winner = newSchedule.ScheduleUserName
		finalResult.WinnerID = newSchedule.ScheduleUserID
		finalResult.Loser = newSchedule.ScheduleOpponent
		finalResult.LoserID = newSchedule.ScheduleOpponentID
	} else if newSchedule.ScheduleResult == "L" {
		finalResult.Loser = newSchedule.ScheduleUserName
		finalResult.LoserID = newSchedule.ScheduleUserID
		finalResult.Winner = newSchedule.ScheduleOpponent
		finalResult.WinnerID = newSchedule.ScheduleOpponentID
	}
	wg.Add(1)
	go c.ScheduleUsecase.updateScheduleBadmintonS(&newSchedule, &finalResult, &wg, &result)
	wg.Wait()
	if result.Err != nil {
		result := model.Response{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: "perbarui jadwal gagal",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)

		w.Write([]byte(byteOfResult))
	} else {
		result := model.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "perbarui jadwal berhasil",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteOfResult))
	}

	finalResult = resultMatch{
		Winner:   "",
		WinnerID: "",
		Loser:    "",
		LoserID:  "",
	}

}
func (c Controller) HandleFetchSchedule(w http.ResponseWriter, r *http.Request) {

	var scheduleList []*schedule

	id := helper.DecodeParams("id", r)

	status := "A"

	err := c.ScheduleUsecase.getScheduleBadmintonS(&scheduleList, id, status)

	if err != nil {
		result := model.Response{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: "ambil jadwal gagal",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)

		w.Write([]byte(byteOfResult))

	} else {

		result := model.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "ambil jadwal berhasil",
			Data:    scheduleList,
		}

		byteOfResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteOfResult))
	}

}
func (c Controller) HandleFetchDetailSchedule(w http.ResponseWriter, r *http.Request) {

	id := helper.DecodeParams("id", r)
	var result model.Result

	c.ScheduleUsecase.getScheduleDetailBadmintonS(&result, &id)

	if result.Err != nil {
		result := model.Response{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: "ambil detail jadwal gagal",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)

		w.Write([]byte(byteOfResult))

	} else {

		result := model.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "ambil detail jadwal berhasil",
			Data:    result.Res,
		}

		byteOfResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteOfResult))
	}

}
func (c Controller) HandleFetchHistory(w http.ResponseWriter, r *http.Request) {

	var scheduleList []*schedule

	id := helper.DecodeParams("id", r)

	status := "I"

	err := c.ScheduleUsecase.getScheduleBadmintonS(&scheduleList, id, status)

	if err != nil {
		result := model.Response{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: "ambil jadwal gagal",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)

		w.Write([]byte(byteOfResult))

	} else {

		result := model.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "ambil jadwal berhasil",
			Data:    scheduleList,
		}

		byteOfResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteOfResult))
	}

}
func (c Controller) HandleFetchAllSchedule(w http.ResponseWriter, r *http.Request) {

	var result model.Result
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	convertPage, _ := strconv.Atoi(page)
	convertLimit, _ := strconv.Atoi(limit)
	page = strconv.Itoa(convertPage*convertLimit - convertLimit)
	
	c.ScheduleUsecase.getScheduleBadmintonSALL(&result, &page, &limit)

	if result.Err != nil {
		result := model.Response{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: "ambil semua jadwal gagal",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)

		w.Write([]byte(byteOfResult))

	} else {

		result := model.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "ambil semua jadwal berhasil",
			Data:    result.Res,
		}

		byteOfResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteOfResult))
	}

}
