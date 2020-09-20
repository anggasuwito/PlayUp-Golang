package user

import (
	"database/sql"
	"log"
	"playUp/main/master/model"
	query "playUp/main/master/utils/db"
	generate "playUp/main/master/utils/generator"
)

type UserRepo struct {
	db *sql.DB
}

type UserRepository interface {
	UserRegister(user *User) error
	UpdateUserProfile(user *User, res *model.Result) error
	UpdateUserProfileOAuth(updateUser *User, res *model.Result, accountType *string) error
	UserLoginOAuth(res *model.Result, user *User) error
	ValidateUserRegister(user *User) (*User, error)
	GetUserList(res *model.Result, page, limit *string) error
	GetUserByID(res *model.Result, id *string) error
	UpdatePhotoProfile(user *User, id, photoPath string) error
}

func NewUserRepo(db *sql.DB) UserRepository {
	return UserRepo{db: db}
}

// UpdateUserProfile -> update user profile

func (d UserRepo) UpdateUserProfile(user *User, res *model.Result) error {
	updatedUser := new(User)

	tx, err := d.db.Begin()

	if err != nil {
		return err
	}

	stmt, _ := tx.Prepare(query.UPDATE_USER_PROFILE)

	_, err = stmt.Exec(user.Username, user.UserFullName, user.UserGender, user.UserEmail, user.ID)

	if err != nil {
		tx.Rollback()
		log.Printf("-> error failed update user profile 1 %v \n", err)
		return err
	}

	err = tx.QueryRow(query.SELECT_USER_BY_ID, user.ID).Scan(&updatedUser.ID, &updatedUser.PhotoProfile, &updatedUser.Username, &updatedUser.UserFullName, &updatedUser.UserGender, &updatedUser.UserEmail, &updatedUser.Password, &updatedUser.Created, &updatedUser.Updated, &updatedUser.FacebookAccount, &updatedUser.GoogleAccount)

	if err != nil {
		tx.Rollback()
		log.Printf("-> error failed update user profile 2 %v \n", err)
		return err
	}

	res.Res = updatedUser
	tx.Commit()
	return nil
}

// UserRegister -> register new user
func (d UserRepo) UserRegister(user *User) error {

	uuidUser := generate.GenerateUUID()
	uuidRank := generate.GenerateUUID()
	tx, err := d.db.Begin()

	if err != nil {
		return err
	}

	stmt, _ := tx.Prepare(query.REGISTER_USER)

	_, err = stmt.Exec(uuidUser, user.PhotoProfile, user.Username, user.UserFullName, user.UserGender, user.UserEmail, user.Password)
	if err != nil {
		log.Printf("error failed register new user %v \n", err)
		tx.Rollback()
		return err

	}
	stmt2, _ := tx.Prepare(query.INSERT_RANK)

	_, err = stmt2.Exec(uuidRank, uuidUser, 0, 0)
	if err != nil {

		tx.Rollback()
		return err

	}
	tx.Commit()
	return nil

}

// ValidateUserRegister -> check wheather user is registered or not
func (d UserRepo) ValidateUserRegister(user *User) (*User, error) {

	regUser := new(User)
	var err error

	if user.UserEmail == "" {
		err = d.db.QueryRow(query.SELECT_USER_BY_USERNAME, user.Username).Scan(&regUser.ID, &regUser.PhotoProfile, &regUser.Username, &regUser.UserFullName, &regUser.UserGender, &regUser.UserEmail, &regUser.Password, &regUser.RankID, &regUser.RankUserMatchCount, &regUser.RankUserMatchGrade)
	} else {
		err = d.db.QueryRow(query.SELECT_USER_BY_USERNAME_OR_EMAIL, user.Username, user.UserEmail).Scan(&regUser.ID, &regUser.PhotoProfile, &regUser.Username, &regUser.UserFullName, &regUser.UserGender, &regUser.UserEmail, &regUser.Password, &regUser.Created, &regUser.Updated)
	}

	if err != nil {
		log.Printf(" -> belum terdaftar %v \n", err)
		return nil, err
	}

	return regUser, nil

}

// GetUserList -> get all user
func (d UserRepo) GetUserList(res *model.Result, page, limit *string) error {

	var userList []*User
	var totalData model.Total
	err := d.db.QueryRow(query.SELECT_USER_TOTAL_COUNT).Scan(&totalData)
	if err != nil {
		log.Printf("error failed get all user 1 %v \n", err)
		return err
	}
	rows, err := d.db.Query(query.SELECT_ALL_USER, *page, *limit)
	if err != nil {
		log.Printf("error failed get all user 2 %v \n", err)
		return err
	}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.PhotoProfile, &user.Username, &user.UserFullName, &user.UserGender, &user.UserEmail, &user.Password, &user.Created, &user.Updated)
		if err != nil {
			log.Printf("error failed get all user 3 %v \n", err)
			return err
		}
		userList = append(userList, &user)
	}
	res.Res = map[string]interface{}{
		"total_data": totalData,
		"result":     userList,
	}
	return nil

}

// GetUserByID -> get invidual user by ID
func (d UserRepo) GetUserByID(res *model.Result, id *string) error {

	user := new(User)

	err := d.db.QueryRow(query.SELECT_USER_BY_USER_ID, id).Scan(&user.ID, &user.PhotoProfile, &user.Username, &user.UserFullName, &user.UserGender, &user.UserEmail, &user.Password, &user.RankID, &user.RankUserMatchCount, &user.RankUserMatchGrade, &user.FacebookAccount, &user.GoogleAccount)
	if err != nil {
		log.Printf("error failed get user by id %v \n", err)
		return err
	}
	res.Res = user

	return nil

}

// UpdatePhotoProfile -> update user photo profile
func (d UserRepo) UpdatePhotoProfile(updatedUser *User, id, photoPath string) error {

	tx, _ := d.db.Begin()

	stmt, err := tx.Prepare(query.UPDATE_USER_PHOTO_PROFILE)

	if err != nil {
		log.Printf("error failed update photo profile 1 %v \n", err)
		return err
	}
	_, err = stmt.Exec(photoPath, id)
	if err != nil {
		tx.Rollback()
		log.Printf("error failed update photo profile 2 %v \n", err)
		return err
	}
	err = tx.QueryRow(query.SELECT_USER_BY_USER_ID, id).Scan(&updatedUser.ID, &updatedUser.PhotoProfile, &updatedUser.Username, &updatedUser.UserFullName, &updatedUser.UserGender, &updatedUser.UserEmail, &updatedUser.Password, &updatedUser.RankID, &updatedUser.RankUserMatchCount, &updatedUser.RankUserMatchGrade, &updatedUser.FacebookAccount, &updatedUser.GoogleAccount)
	if err != nil {
		log.Printf("error failed update photo profile 3 %v \n", err)
		tx.Rollback()
		return err

	}
	tx.Commit()
	return nil

}

// UpdatePhotoProfile -> update user photo profile
func (d UserRepo) UpdateUserProfileOAuth(updateUser *User, res *model.Result, accountType *string) error {
	tx, _ := d.db.Begin()
	stmt1, err := tx.Prepare(`update master_user set user_facebook_account=? where user_id=?`)

	if err != nil {
		log.Printf("gagal query 1 %v", err)
		return err
	}
	stmt2, err := tx.Prepare(`update master_user set user_google_account=? where user_id=?`)
	if err != nil {
		log.Printf("gagal query 2 %v", err)
		return err
	}
	if *accountType == "google" {
		_, err := stmt2.Exec(updateUser.GoogleAccount, updateUser.ID)
		if err != nil {
			tx.Rollback()
			log.Println(err)
			return err
		}
	} else {
		_, err := stmt1.Exec(updateUser.FacebookAccount, updateUser.ID)
		if err != nil {
			tx.Rollback()
			log.Println(err)
			return err
		}
	}
	tx.Commit()
	return nil
}
func (d UserRepo) UserLoginOAuth(res *model.Result, user *User) error {
	tx, _ := d.db.Begin()
	loginUser := new(User)
	if user.FacebookAccount == "" {
		err := tx.QueryRow(query.SELECT_USER_BY_GOOGLE, user.GoogleAccount).Scan(&loginUser.ID, &loginUser.PhotoProfile, &loginUser.Username, &loginUser.UserFullName, &loginUser.UserGender, &loginUser.UserEmail, &loginUser.Password, &loginUser.GoogleAccount, &loginUser.RankID, &loginUser.RankUserMatchCount, &loginUser.RankUserMatchGrade)
		if err != nil {
			log.Printf("Error login OAuth google %v", err)
			return err
		}
	} else {
		err := tx.QueryRow(query.SELECT_USER_BY_FACEBOOK, user.FacebookAccount).Scan(&loginUser.ID, &loginUser.PhotoProfile, &loginUser.Username, &loginUser.UserFullName, &loginUser.UserGender, &loginUser.UserEmail, &loginUser.Password, &loginUser.GoogleAccount, &loginUser.RankID, &loginUser.RankUserMatchCount, &loginUser.RankUserMatchGrade)
		if err != nil {
			log.Printf("Error login OAuth facebook %v", err)
			return err
		}
	}

	tx.Commit()
	res.Res = loginUser
	return nil
}
