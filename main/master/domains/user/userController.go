package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"playUp/main/master/model"
	"playUp/main/master/utils/helper"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

type Controller struct {
	UserUsecase UserUsecaseInterface
}

func NewController(db *sql.DB) *Controller {
	return &Controller{UserUsecase: NewUserUsecase(db)}
}

func (c Controller) UserLogin(w http.ResponseWriter, r *http.Request) {

	var newUser User
	var result model.Result
	var wg sync.WaitGroup
	_ = json.NewDecoder(r.Body).Decode(&newUser)

	wg.Add(1)
	go c.UserUsecase.UserLogin(&newUser, &result, &wg)
	wg.Wait()
	if result.Err != nil {
		result := model.Response{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: "username atau password tidak cocok",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)

		w.Write([]byte(byteOfResult))
	} else {
		result := model.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Login berhasil",
			Data:    result.Res,
		}
		byteOfResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteOfResult))
	}
}

// UserRegister -> register new user
func (c Controller) UserRegister(w http.ResponseWriter, r *http.Request) {

	var newUser User
	var result model.Result
	var wg sync.WaitGroup

	_ = json.NewDecoder(r.Body).Decode(&newUser)

	wg.Add(1)
	go c.UserUsecase.UserRegister(&newUser, &result, &wg)
	wg.Wait()

	if result.Err != nil {
		result := model.Response{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: "Username atau email sudah terdaftar",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)

		w.Write([]byte(byteOfResult))
	} else {
		result := model.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Daftar user baru berhasil",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteOfResult))
	}

}

// UserUpdate -> update user profile
func (c Controller) UserUpdate(w http.ResponseWriter, r *http.Request) {

	var updatedUser User
	var result model.Result
	var wg sync.WaitGroup

	_ = json.NewDecoder(r.Body).Decode(&updatedUser)

	wg.Add(1)
	go c.UserUsecase.UpdateUserProfile(&updatedUser, &result, &wg)
	wg.Wait()

	if result.Err != nil {
		result := model.Response{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: "Perbarui profile user gagal",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)

		w.Write([]byte(byteOfResult))
	} else {
		result := model.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "perbarui profile user berhasil",
			Data:    result.Res,
		}
		byteOfResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteOfResult))
	}

}

func (c Controller) GetUserList(w http.ResponseWriter, r *http.Request) {

	var result model.Result
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	convertPage, _ := strconv.Atoi(page)
	convertLimit, _ := strconv.Atoi(limit)
	page = strconv.Itoa(convertPage*convertLimit - convertLimit)

	c.UserUsecase.GetUserList(&result, &page, &limit)

	if result.Err != nil {
		result := model.Response{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: "Ambil semua user gagal",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)

		w.Write([]byte(byteOfResult))
	} else {
		result := model.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Ambil semua user berhasil",
			Data:    result.Res,
		}
		byteOfResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteOfResult))
	}

}
func (c Controller) GetUserByID(w http.ResponseWriter, r *http.Request) {

	var result model.Result
	var wg sync.WaitGroup
	id := helper.DecodeParams("id", r)

	wg.Add(1)
	c.UserUsecase.GetUserByID(&result, &id, &wg)
	wg.Wait()
	if result.Err != nil {
		result := model.Response{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: "Ambil user berdasarkan id gagal",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)

		w.Write([]byte(byteOfResult))
	} else {
		result := model.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Ambil user berdasarkan id berhasil",
			Data:    result.Res,
		}
		byteOfResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteOfResult))
	}

}

//PostImage AddImage
func (c Controller) PostImage(w http.ResponseWriter, r *http.Request) {

	//membatasi max upload 1024
	r.ParseMultipartForm(1024)
	id := r.FormValue("id")
	var updatedUser User
	log.Println(id)
	//mengambil key dari file dari formdata
	image, handlerImage, err := r.FormFile("image")
	if err != nil {
		log.Println(`Error while parsing file`, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	fmt.Println("image = ", image)
	fmt.Println("handler = ", handlerImage)
	fmt.Println("handler filename = ", handlerImage.Filename)
	fmt.Println("handler header = ", handlerImage.Header)
	fmt.Println("handler size = ", handlerImage.Size)

	//mengambil key dari text dari formdata
	var imageModel model.Image
	data := r.FormValue("data")
	_ = json.Unmarshal([]byte(data), &imageModel)
	fmt.Println("imagemodel username = ", imageModel.Username)
	fmt.Println("imagemodel image name = ", imageModel.ImageName)

	//size adalah int64
	fmt.Printf("handlerSizeOriginalType = %v %T \n", handlerImage.Size, handlerImage.Size)

	//add username buat prefix new photo name
	newPhotoName := helper.CopyImageAndRename(image, handlerImage, imageModel.Username)

	//Rapiin, buat use case dan reponya
	//bikin query update user set photo = newPhotoName(line 143) where username = imageModel.Username (line 134 137)
	//buat response balikan pake wrapper yang dimana data isinya data user yang sudah diubah, karna akan di ambil kembali newPhotoName kedepan

	//convert int64 to int to String
	// imageSize := strconv.Itoa(int(handlerImage.Size))

	err = c.UserUsecase.UpdatePhotoProfile(&updatedUser, id, newPhotoName)

	if err != nil {
		result := model.Response{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: "update photo profile gagal",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)

		w.Write([]byte(byteOfResult))
	} else {
		result := model.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "update berhasil",
			Data:    updatedUser,
		}
		byteOfResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteOfResult))
	}
	//hapus line 153 ganti dengan response terbaru
	// response := "old image name = " + handlerImage.Filename + "\nold image size = " + imageSize + "\nnew photo name = " + newPhotoName
	// //terakhir hapus semua fmt.println yang ada pada controller post-image agar bersih

	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte(response))
}

//GetImage AddImage
func (c Controller) GetImage(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	idPhoto := param["id"]
	dir, _ := os.Getwd()
	fileLocationWithDir := filepath.Join(dir, "uploads", idPhoto)
	w.Header().Set("Content-Type", "image/jpeg")
	http.ServeFile(w, r, fileLocationWithDir)
}

// UserUpdateGoogle -> update user google profile
func (c Controller) UserUpdateGoogle(w http.ResponseWriter, r *http.Request) {
	var accountType = "google"
	var updatedUser User
	var result model.Result
	var wg sync.WaitGroup

	_ = json.NewDecoder(r.Body).Decode(&updatedUser)

	wg.Add(1)
	go c.UserUsecase.UpdateUserProfileOAuth(&result,&updatedUser, &wg,&accountType)
	wg.Wait()

	if result.Err != nil {
		result := model.Response{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: "Perbarui profile user google gagal",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)

		w.Write([]byte(byteOfResult))
	} else {
		result := model.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "perbarui profile user google berhasil",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteOfResult))
	}

}

// UserUpdateFacebook -> update user facebook profile
func (c Controller) UserUpdateFacebook(w http.ResponseWriter, r *http.Request) {

	var accountType = "facebook"
	var updatedUser User
	var result model.Result
	var wg sync.WaitGroup

	_ = json.NewDecoder(r.Body).Decode(&updatedUser)

	wg.Add(1)
	go c.UserUsecase.UpdateUserProfileOAuth(&result,&updatedUser, &wg,&accountType)
	wg.Wait()

	if result.Err != nil {
		result := model.Response{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: "Perbarui profile user facebook gagal",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)

		w.Write([]byte(byteOfResult))
	} else {
		result := model.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "perbarui profile user facebook berhasil",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteOfResult))
	}

}
func (c Controller) UserLoginOAuth(w http.ResponseWriter, r *http.Request) {

	var newUser User
	var result model.Result
	var wg sync.WaitGroup
	_ = json.NewDecoder(r.Body).Decode(&newUser)

	wg.Add(1)
	go c.UserUsecase.UserLoginOAuth(&newUser, &result, &wg)
	wg.Wait()
	if result.Err != nil {
		result := model.Response{
			Status:  "failed",
			Code:    http.StatusInternalServerError,
			Message: "username atau password tidak cocok",
			Data:    nil,
		}
		byteOfResult, _ := json.Marshal(result)

		w.Write([]byte(byteOfResult))
	} else {
		result := model.Response{
			Status:  "success",
			Code:    http.StatusOK,
			Message: "Login berhasil",
			Data:    result.Res,
		}
		byteOfResult, _ := json.Marshal(result)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteOfResult))
	}
}
