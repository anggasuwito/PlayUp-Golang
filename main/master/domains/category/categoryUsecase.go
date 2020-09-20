package category

import (
	"database/sql"
	"errors"
	"log"
)

type CategoryUsecase struct {
	db       *sql.DB
	CategoryRepo CategoryRepository
}

func (c CategoryUsecase) FindCategory(id string) (Category, error) {
	findCategory, err := c.CategoryRepo.FindCategory(id)
	if err != nil {
		log.Println("ini error", err)
		return findCategory, errors.New("CategoryID NOY FOUND")
	}

	return findCategory, nil
}

type CategoryUsecaseInterface interface {
	PostCategory(category *Category) error
	GetCategoryList(categoryList *[]*Category) error
	PutCategory(category *Category) error
	DeleteCategory(id string) error
	FindCategory(string) (Category, error)
	ValidateCategory(category *Category) (*Category, error)
}

func NewCategoryUsecase(db *sql.DB) CategoryUsecaseInterface {
	return CategoryUsecase{db, NewCategoryRepo(db)}
}

func (c CategoryUsecase) ValidateCategory(category *Category) (*Category, error) {

	registeredUser, err := c.CategoryRepo.ValidateCategory(category)

	if err != nil {
		log.Println("ini error", err)
		return nil, errors.New("CategoryID NOY FOUND")
	}

	return registeredUser, nil
}

func (c CategoryUsecase) PostCategory(category *Category) error {

	addCategory, err := c.CategoryRepo.ValidateCategory(category)
	log.Println(category.CategoryName)

	if err != nil {
		err := c.CategoryRepo.PostCategory(category)

		if err != nil {
			return err
		}

		return nil
	}

	if addCategory.CategoryName == category.CategoryName {

		log.Println("Sudah terdaftar")
		return errors.New("Category sudah terdaftar")

	}
	return nil
}

func (c CategoryUsecase) PutCategory(category *Category) error {
	err := c.CategoryRepo.PutCategory(category)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (c CategoryUsecase) DeleteCategory(id string) error {
	err := c.CategoryRepo.DeleteCategory(id)
	if err != nil {
		return err
	}
	return nil
}

func (c CategoryUsecase) GetCategoryList(categoryList *[]*Category) error {
	err := c.CategoryRepo.GetCategoryList(categoryList)
	if err != nil {
		return err
	}
	return nil
}
