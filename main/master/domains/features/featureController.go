package features

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"playUp/main/master/model"
	"playUp/main/master/utils/helper"
)

type Controller struct {
	FeatureUsecase FeatureUsecaseInterface
}

func NewController(db *sql.DB) *Controller {
	return &Controller{FeatureUsecase: NewFeatureUsecase(db)}
}

func (c Controller) FindFeature(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	id := helper.DecodeParams("id", r)
	result, err := c.FeatureUsecase.FindFeature(id)
	if err != nil {
		cek := model.Response{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: "Cari Feature gagal",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(cek)

		w.Write([]byte(byteOfResult))
	} else {
		cek := model.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Cari Feature berhasil",
			Data:    result,
		}
		byteOfResult, _ := json.Marshal(cek)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteOfResult))
	}
}

func (c Controller) GetFeatureList(w http.ResponseWriter, r *http.Request) {

	var featureList []*Feature

	err := c.FeatureUsecase.GetFeatureList(&featureList)
	if err != nil {
		result := model.Response{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: "Ambil semua feature gagal",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)

		w.Write([]byte(byteOfResult))
	} else {
		result := model.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Ambil semua feature berhasil",
			Data:    featureList,
		}
		byteOfResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteOfResult))
	}

}

func (c Controller) PostFeature(w http.ResponseWriter, r *http.Request) {

	var newFeature *Feature

	_ = json.NewDecoder(r.Body).Decode(&newFeature)

	err := c.FeatureUsecase.PostFeature(newFeature)

	if err != nil {
		result := model.Response{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: "Feature sudah ada",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)

		w.Write([]byte(byteOfResult))
	} else {
		result := model.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Daftar feature baru berhasil",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteOfResult))
	}

}

func (c Controller) PutFeature(w http.ResponseWriter, r *http.Request) {

	newFeature := new(Feature)

	_ = json.NewDecoder(r.Body).Decode(&newFeature)

	err := c.FeatureUsecase.PutFeature(newFeature)
	if err != nil {
		result := model.Response{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: "Perbaharui feature gagal",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)

		w.Write([]byte(byteOfResult))
	} else {
		result := model.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Perbaharui feature berhasil",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteOfResult))
	}

}

func (c Controller) DeleteFeature(w http.ResponseWriter, r *http.Request) {

	id := helper.DecodeParams("id", r)

	err := c.FeatureUsecase.DeleteFeature(id)
	if err != nil {
		result := model.Response{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: "delete feature gagal",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)

		w.Write([]byte(byteOfResult))
	} else {
		result := model.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "delete feature berhasil",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteOfResult))
	}

}
