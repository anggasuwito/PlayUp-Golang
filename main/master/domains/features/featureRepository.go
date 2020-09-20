package features

import (
	"database/sql"
	"errors"
	"log"
	query "playUp/main/master/utils/db"
	generate "playUp/main/master/utils/generator"
)

type FeatureRepo struct {
	db *sql.DB
}
// FindFeature -> find feature
func (c FeatureRepo) FindFeature(id string) (Feature, error) {

	var feature Feature
	stmt := query.SELECT_FEATURE_BY_ID
	err := c.db.QueryRow(stmt, id).Scan(&feature.ID, &feature.FeatureName, &feature.FeatureDescription, &feature.FeatureImage, &feature.FeatureStatus, &feature.Created, &feature.Updated)
	if err != nil {
		return feature, err
	}
	return feature, nil
}
// ValidateFeature -> check wheater feature is registered or not
func (c FeatureRepo) ValidateFeature(feature *Feature) (*Feature, error) {
	regUser := new(Feature)
	var err error

	if feature.ID == "" {
		err = c.db.QueryRow(query.SELECT_FEATURE_BY_NAME, feature.FeatureName).Scan(&regUser.ID, &regUser.FeatureName, &regUser.FeatureDescription, &regUser.FeatureImage, &regUser.FeatureStatus, &regUser.Created, &regUser.Updated)
	}

	if err != nil {
		log.Println(err)
		return nil, errors.New("FAILED FIND")
	}

	return regUser, nil
}
// PostFeature -> create new feature
func (c FeatureRepo) PostFeature(feature *Feature) error {
	uuid := generate.GenerateUUID()
	tx, err := c.db.Begin()
	if err != nil {
		log.Println(err)
		log.Println("DATA NOT EXIST")
		return err
	}
	stmt, _ := tx.Prepare(query.INSERT_FEATURE)

	_, err = stmt.Exec(uuid, feature.FeatureName, feature.FeatureDescription, feature.FeatureImage)
	if err != nil {
		log.Println(err)
		tx.Rollback()

	}
	tx.Commit()
	return nil
}
// PutFeature -> change existing feature
func (c FeatureRepo) PutFeature(feature *Feature) error {
	tx, err := c.db.Begin()

	if err != nil {
		log.Println(err)
	}
	stmt, _ := tx.Prepare(query.UPDATE_FEATURE)

	_, err = stmt.Exec(feature.FeatureName, feature.FeatureDescription, feature.FeatureImage, feature.ID)

	if err != nil {
		log.Println(err)
		tx.Rollback()
	}
	tx.Commit()
	return nil
}
// DeleteFeature -> Delete Feature
func (c FeatureRepo) DeleteFeature(id string) error {
	tx, err := c.db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}
	stmt, _ := tx.Prepare(query.DELETE_FEATURE)

	_, err = stmt.Exec(id)
	if err != nil {
		tx.Rollback()
		return err
	}

	stmt.Close()
	tx.Commit()
	return nil
}
// GetFeatureList -> get feature list
func (c FeatureRepo) GetFeatureList(featureList *[]*Feature) error {

	rows, err := c.db.Query(query.SELECT_ALL_FEATURE)
	if err != nil {
		return err
	}
	for rows.Next() {
		var feature Feature
		err := rows.Scan(&feature.ID, &feature.FeatureName, &feature.FeatureDescription, &feature.FeatureImage, &feature.FeatureStatus, &feature.Created, &feature.Updated)
		if err != nil {
			return err
		}
		*featureList = append(*featureList, &feature)
	}
	return nil
}

type FeatureRepository interface {
	PostFeature(feature *Feature) error
	PutFeature(feature *Feature) error
	DeleteFeature(id string) error
	FindFeature(string) (Feature, error)
	GetFeatureList(featureList *[]*Feature) error
	ValidateFeature(feature *Feature) (*Feature, error)
}

func NewFeatureRepo(db *sql.DB) FeatureRepository {
	return FeatureRepo{db: db}
}
