package category

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"playUp/main/master/model"
	"playUp/main/master/utils/helper"
)

type Controller struct {
	CategoryUsecase CategoryUsecaseInterface
}

func NewController(db *sql.DB) *Controller {
	return &Controller{CategoryUsecase: NewCategoryUsecase(db)}
}

func (c Controller) FindCategory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	id := helper.DecodeParams("id", r)
	result, err := c.CategoryUsecase.FindCategory(id)
	if err != nil {
		cek := model.Response{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: "Cari category gagal",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(cek)

		w.Write([]byte(byteOfResult))
	} else {
		cek := model.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Cari category berhasil",
			Data:    result,
		}
		byteOfResult, _ := json.Marshal(cek)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteOfResult))
	}
}

func (c Controller) GetCategoryList(w http.ResponseWriter, r *http.Request) {

	var categoryList []*Category

	err := c.CategoryUsecase.GetCategoryList(&categoryList)
	if err != nil {
		result := model.Response{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: "Ambil semua category gagal",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)

		w.Write([]byte(byteOfResult))
	} else {
		result := model.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Ambil semua category berhasil",
			Data:    categoryList,
		}
		byteOfResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteOfResult))
	}

}

func (c Controller) PostCategory(w http.ResponseWriter, r *http.Request) {

	var newCategpry *Category

	_ = json.NewDecoder(r.Body).Decode(&newCategpry)

	err := c.CategoryUsecase.PostCategory(newCategpry)

	if err != nil {
		result := model.Response{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: "Category sudah ada",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)

		w.Write([]byte(byteOfResult))
	} else {
		result := model.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Daftar category baru berhasil",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteOfResult))
	}

}

func (c Controller) PutCategory(w http.ResponseWriter, r *http.Request) {

	newCategory := new(Category)

	_ = json.NewDecoder(r.Body).Decode(&newCategory)

	err := c.CategoryUsecase.PutCategory(newCategory)
	if err != nil {
		result := model.Response{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: "perbarui category gagal",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)

		w.Write([]byte(byteOfResult))
	} else {
		result := model.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "perbarui category berhasil",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteOfResult))
	}

}

func (c Controller) DeleteCategory(w http.ResponseWriter, r *http.Request) {

	id := helper.DecodeParams("id", r)

	err := c.CategoryUsecase.DeleteCategory(id)
	if err != nil {
		result := model.Response{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: "delete category gagal",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)

		w.Write([]byte(byteOfResult))
	} else {
		result := model.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "delete category berhasil",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteOfResult))
	}

}
