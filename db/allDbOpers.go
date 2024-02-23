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

func (r *Repository) UpdateProfile(profiles *models.CommissionProfiles) (*models.CommissionProfiles, error) {
	err := r.Db.Model(&models.CommissionProfiles{}).Update(&profiles).Scan(&profiles).Error
	if err != nil {
		return nil, err
	}
	return profiles, nil
}

func (r *Repository) DeleteProfile(profiles models.CommissionProfiles) error {
	if err := r.Db.Model(&models.CommissionProfiles{}).Where("id = ?", profiles.Id).UpdateColumns(map[string]interface{}{
		"active":     false,
		"updated_at": profiles.UpdatedAt,
		"updated_by": profiles.UpdatedBy,
	}).Error; err != nil {
		return err
	}
	return nil
}
