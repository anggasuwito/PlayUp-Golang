package features

import (
	"database/sql"
	"errors"
	"log"
)

type FeatureUsecase struct {
	db       *sql.DB
	FeatureRepo FeatureRepository
}

type FeatureUsecaseInterface interface {
	PostFeature(feature *Feature) error
	PutFeature(feature *Feature) error
	DeleteFeature(id string) error
	FindFeature(string) (Feature, error)
	GetFeatureList(featureList *[]*Feature) error
	ValidateFeature(feature *Feature) (*Feature, error)
}

func NewFeatureUsecase(db *sql.DB) FeatureUsecaseInterface {
	return FeatureUsecase{db, NewFeatureRepo(db)}
}

func (c FeatureUsecase) FindFeature(id string) (Feature, error) {
	findFeature, err := c.FeatureRepo.FindFeature(id)
	if err != nil {
		log.Println("ini error", err)
		return findFeature, errors.New("FeatureID NOY FOUND")
	}

	return findFeature, nil
}

func (c FeatureUsecase) ValidateFeature(feature *Feature) (*Feature, error) {

	registeredUser, err := c.FeatureRepo.ValidateFeature(feature)

	if err != nil {
		log.Println("ini error", err)
		return nil, errors.New("ID NOT FOUND")
	}

	return registeredUser, nil
}

func (c FeatureUsecase) PostFeature(feature *Feature) error {

	addCategory, err := c.FeatureRepo.ValidateFeature(feature)
	log.Println(feature.FeatureName)

	if err != nil {
		err := c.FeatureRepo.PostFeature(feature)

		if err != nil {
			return err
		}

		return nil
	}

	if addCategory.FeatureName == feature.FeatureName {

		log.Println("Sudah terdaftar")
		return errors.New("Feature already exist")

	}
	return nil
}

func (c FeatureUsecase) PutFeature(feature *Feature) error {
	err := c.FeatureRepo.PutFeature(feature)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (c FeatureUsecase) DeleteFeature(id string) error {
	err := c.FeatureRepo.DeleteFeature(id)
	if err != nil {
		return err
	}
	return nil
}

func (c FeatureUsecase) GetFeatureList(featureList *[]*Feature) error {
	err := c.FeatureRepo.GetFeatureList(featureList)
	if err != nil {
		return err
	}
	return nil
}
