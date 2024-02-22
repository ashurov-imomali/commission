package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"main/models"
)

type Repository struct {
	Db *gorm.DB
}

func GetConnection(dbSetting string) (*Repository, error) {
	db, err := gorm.Open("postgres", dbSetting)
	if err != nil {
		return nil, err
	}
	return &Repository{Db: db}, nil
}

func (r *Repository) CreateProfile(profiles *models.CommissionProfiles) (int64, error) {
	if err := r.Db.Create(&profiles).Error; err != nil {
		return 0, nil
	}
	return profiles.Id, nil
}

func (r *Repository) CreateRules(rules []models.CommissionRules, profileId int64) error {
	for _, rule := range rules {
		rule.Active = true
		rule.ProfileId = profileId
		err := r.Db.Create(&rule).Error
		if err != nil {
			return err
		}
	}
	return nil
}
