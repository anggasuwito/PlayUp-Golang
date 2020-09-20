package category

import (
	"database/sql"
	"errors"
	"log"
	query "playUp/main/master/utils/db"
	generate "playUp/main/master/utils/generator"
)

type CategoryRepo struct {
	db *sql.DB
}

func (c CategoryRepo) FindCategory(id string) (Category, error) {
	var category Category
	stmt := query.SELECT_CATEGORY_BY_ID
	err := c.db.QueryRow(stmt, id).Scan(&category.ID, &category.CategoryName, &category.CategoryImage, &category.CategoryStatus, &category.Created, &category.Updated)
	if err != nil {
		return category, err
	}
	return category, nil
}

func (c CategoryRepo) ValidateCategory(category *Category) (*Category, error) {
	regUser := new(Category)
	var err error

	if category.ID == "" {
		err = c.db.QueryRow(query.SELECT_CATEGORY_BY_NAME, category.CategoryName).Scan(&regUser.ID, &regUser.CategoryName, &regUser.CategoryImage, &regUser.CategoryStatus, &regUser.Created, &regUser.Updated)
	}

	if err != nil {
		log.Println(err)
		return nil, errors.New("FAILED FIND")
	}

	return regUser, nil
}

func (c CategoryRepo) PostCategory(category *Category) error {

	uuid := generate.GenerateUUID()
	tx, err := c.db.Begin()
	if err != nil {
		log.Println(err)
		log.Println("DATA NOT EXIST")
		return err
	}
	stmt, _ := tx.Prepare(query.INSERT_CATEGORY)

	_, err = stmt.Exec(uuid, category.CategoryName, category.CategoryImage)
	if err != nil {
		log.Println(err)
		tx.Rollback()

	}
	tx.Commit()
	return nil
}

func (c CategoryRepo) PutCategory(category *Category) error {
	tx, err := c.db.Begin()

	if err != nil {
		log.Println(err)
	}
	stmt, _ := tx.Prepare(query.UPDATE_CATEGORY)

	_, err = stmt.Exec(category.CategoryName,category.CategoryImage, category.ID)

	if err != nil {
		log.Println("GAGAL UPDATE")
		log.Println(err)
		tx.Rollback()
	}
	tx.Commit()
	return nil
}

func (c CategoryRepo) DeleteCategory(id string) error {
	tx, err := c.db.Begin()
	if err != nil {
		log.Println(err)
	}
	stmt, _ := tx.Prepare(query.DELETE_CATEGORY)

	_, err = stmt.Exec(id)
	if err != nil {
		return tx.Rollback()
	}

	stmt.Close()
	return tx.Commit()
}

func (c CategoryRepo) GetCategoryList(categoryList *[]*Category) error {
	
	rows, err := c.db.Query(query.SELECT_ALL_CATEGORY)
	if err != nil {
		return err
	}
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.ID, &category.CategoryName, &category.CategoryImage, &category.CategoryStatus, &category.Created, &category.Updated)
		if err != nil {
			return err
		}
		*categoryList = append(*categoryList, &category)
	}
	return nil
}

type CategoryRepository interface {
	PostCategory(category *Category) error
	PutCategory(category *Category) error
	DeleteCategory(id string) error
	FindCategory(string) (Category, error)
	GetCategoryList(categoryList *[]*Category) error
	ValidateCategory(category *Category) (*Category, error)
}

func NewCategoryRepo(db *sql.DB) CategoryRepository {
	return CategoryRepo{db: db}
}
