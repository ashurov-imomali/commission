package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"main/models"
	"time"
)

type Repository struct {
	Db *gorm.DB
}

func GetConnection(dbSetting string) (*Repository, error) {
	db, err := gorm.Open("postgres", dbSetting)
	if err != nil {
		return nil, err
	}
	//err = db.AutoMigrate(&models.TUser{}).Error
	if err != nil {
		return nil, err
	}
	return &Repository{Db: db}, nil
}

func (r *Repository) CreateProfile(profiles *models.CommissionProfiles) (int64, error) {
	//active отправляется
	if err := r.Db.Create(&profiles).Error; err != nil {
		return 0, nil
	}
	return profiles.Id, nil
}

func (r *Repository) CreateRules(rules []models.CommissionRules, profileId int64) error {
	for _, rule := range rules {
		*rule.Active = true
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
		"deleted_at": profiles.UpdatedAt,
		"updated_by": profiles.UpdatedBy,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateRules(rule *models.CommissionRules) (*models.CommissionRules, error) {
	if err := r.Db.Model(&models.CommissionRules{}).Where("id = ?", rule.Id).
		Update(&rule).Scan(&rule).Error; err != nil {
		return nil, err
	}
	return rule, nil
}

func (r *Repository) DeleteRule(rules *models.CommissionRules) error {
	if err := r.Db.Model(&models.CommissionRules{}).Where("id = ?", rules.Id).
		UpdateColumns(map[string]interface{}{
			"active":     false,
			"deleted_at": time.Now(),
		}).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateProfileRules(profileId, updaterID int64) error {
	if err := r.Db.Model(&models.CommissionProfiles{}).Where("id = ?", profileId).UpdateColumns(map[string]interface{}{
		"updated_at": time.Now(),
		"updated_by": updaterID,
	}).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetProfiles(profileName string) ([]models.ProfileResponse, error) {
	joins := r.Db.Table("tcomission_profiles t").
		Select("t.*, uc.fullname as creater_name, uu.fullname as updater_name").
		Joins("join tusers uc on uc.user_id = t.created_by left join tusers uu on uu.user_id = t.updated_by")

	if profileName != "" {
		joins = joins.Where("name like ? and active = ?", "%"+profileName+"%", true)
	}
	var response []models.ProfileResponse
	if err := joins.Scan(&response).Error; err != nil {
		return nil, err
	}
	return response, nil
}

func (r *Repository) GetRules(profId int64) ([]models.CommissionRules, error) {
	var rules []models.CommissionRules
	if err := r.Db.Model(&models.CommissionRules{}).Where("profile_id = ?", profId).
		Find(&rules).Error; err != nil {
		return nil, err
	}
	return rules, nil
}
