package user

import (
	"database/sql"
	"errors"
	"log"
	"playUp/main/master/model"
	password "playUp/main/master/utils/password"
	token "playUp/main/master/utils/token"
	"sync"
)

// UserUsecase -> struct
type UserUsecase struct {
	db       *sql.DB
	UserRepo UserRepository
}

// UserUsecaseInterface -> all user usecase interface
type UserUsecaseInterface interface {
	UserRegister(user *User, res *model.Result, wg *sync.WaitGroup)
	UserLogin(user *User, res *model.Result, wg *sync.WaitGroup)
	UserLoginOAuth(user *User, res *model.Result, wg *sync.WaitGroup)
	UpdateUserProfile(user *User, res *model.Result, wg *sync.WaitGroup)
	GetUserByID(res *model.Result, id *string,wg *sync.WaitGroup)

	UpdateUserProfileOAuth(res *model.Result, updateUser *User,wg *sync.WaitGroup,accountType *string)


	GetUserList(res *model.Result, page, limit *string)
	UpdatePhotoProfile(updatedUser *User, id, photoPath string) error
}

func NewUserUsecase(db *sql.DB) UserUsecaseInterface {
	return UserUsecase{db, NewUserRepo(db)}
}

// UserLogin -> handle login user
func (s UserUsecase) UserLogin(user *User, res *model.Result, wg *sync.WaitGroup) {
	defer wg.Done()

	registeredUser, err := s.UserRepo.ValidateUserRegister(user)

	if registeredUser == nil || err != nil {
		res.Err = errors.New("Username atau password tidak cocok")
	} else {
		result := password.CheckPasswordHash(user.Password, registeredUser.Password)

		if result == true && user.Username == "admin" && user.Username == registeredUser.Username {
			tokenString := token.GenerateToken(user.Username, 600)
			res.Res = map[string]interface{}{
				"token": tokenString,
				"user":  registeredUser,
			}
		} else if result && user.Username == registeredUser.Username {
			res.Res = registeredUser
		} else {
			log.Println("Username atau Password tidak cocok")
			res.Res = nil
			res.Err = errors.New("Username atau password tidak cocok")
		}
	}

}
// UserLogin -> handle login user
func (s UserUsecase) UserLoginOAuth(user *User, res *model.Result, wg *sync.WaitGroup) {
	defer wg.Done()

	err := s.UserRepo.UserLoginOAuth(res,user)
	
	if err != nil {
		res.Err = err
	}

}

// UserRegister -> register new user
func (s UserUsecase) UserRegister(user *User, res *model.Result, wg *sync.WaitGroup) {
	defer wg.Done()

	user.Password, _ = password.HashPassword(user.Password)

	registeredUser, err := s.UserRepo.ValidateUserRegister(user)
	if registeredUser == nil || err != nil {

		err2 := s.UserRepo.UserRegister(user)

		if err2 != nil {
			res.Err = err2
		}

	} else if registeredUser.Username == user.Username || registeredUser.UserEmail == user.UserEmail {
		log.Println("Sudah terdaftar")
		res.Err = errors.New("Username atau Email sudah terdaftar")
	}

}

// UpdateUserProfile -> update user profile
func (s UserUsecase) UpdateUserProfile(user *User, res *model.Result, wg *sync.WaitGroup) {

	defer wg.Done()

	err := s.UserRepo.UpdateUserProfile(user, res)

	if err != nil {
		res.Err = err
	}

}
func (s UserUsecase) GetUserList(res *model.Result, page, limit *string) {

	err := s.UserRepo.GetUserList(res, page, limit)

	if err != nil {
		res.Err = err
	}

}

// GetUserByID -> get user by id
func (s UserUsecase) GetUserByID(res *model.Result, id *string,wg *sync.WaitGroup) {
	defer wg.Done()
	
	err := s.UserRepo.GetUserByID(res, id)

	if err != nil {
		res.Err = err
	}

}

// UpdatePhotoProfile -> update user photo profile
func (s UserUsecase) UpdatePhotoProfile(updatedUser *User, id, photoPath string) error {

	err := s.UserRepo.UpdatePhotoProfile(updatedUser, id, photoPath)

	if err != nil {
		return err
	}

	return nil
}
// UpdatePhotoProfile -> update user photo profile
func (s UserUsecase) UpdateUserProfileOAuth(res *model.Result, updateUser *User,wg *sync.WaitGroup,accountType *string) {
	defer wg.Done()
	err := s.UserRepo.UpdateUserProfileOAuth(updateUser,res,accountType)

	if err != nil {
		res.Err = err
	}

}
