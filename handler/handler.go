package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/db"
	"main/models"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	Repos *db.Repository
}

func GetHandler(repository *db.Repository) *Handler {
	return &Handler{repository}
}

func (h *Handler) Test(c *gin.Context) {
	c.String(200, "ok")
}

func (h *Handler) CreateCommissionProfile(c *gin.Context) {
	var request models.ProfileCreatRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Can't parse to json. Error:" + err.Error()})
		return
	}
	request.Profile.CreatedBy = 2
	err = h.Repos.CreateProfileAndRules(&request)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't create profile and rules"})
		return
	}
	//profileId, err := CreateProfile(&request.Profile)
	//if err != nil {
	//	log.Println(err)
	//	c.JSON(http.StatusInternalServerError, gin.H{"message": "can't create profile"})
	//	return
	//}
	//err = CreateRules(request.Rules, profileId)
	//if err != nil {
	//	log.Println(err)
	//	c.JSON(http.StatusInternalServerError, gin.H{"message": "can't add new rules"})
	//	return
	//}
	c.Status(http.StatusCreated)
}

func (h *Handler) UpdateCommissionProfiles(c *gin.Context) {
	var updProfile models.ProfileCreatRequest
	err := c.ShouldBindJSON(&updProfile)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "can't parse to json"})
		return
	}
	user := 2
	now := time.Now()
	updProfile.Profile.UpdatedBy = int64(user)
	updProfile.Profile.UpdatedAt = &now
	log.Println(updProfile.Rules)
	if updProfile.Profile.Active == false {
		updProfile.Profile.DeletedAt = &now
		err := h.Repos.DeleteProfileAndRules(&updProfile)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "can't delete profile"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
		return
	}
	err = h.Repos.UpdateProfileAndRules(&updProfile)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't update profile"})
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) UpdateCommissionRules(c *gin.Context) {
	var rules models.CommissionRules
	err := c.ShouldBindJSON(&rules)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "can't parse 2 json"})
		return
	}
	var userId int64 = 2
	now := time.Now()
	err = h.Repos.UpdateProfileRules(rules.ProfileId, userId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't update profile"})
		return
	}
	if rules.Active == false {
		if err := h.Repos.DeleteRule(&rules); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "can't delete rule"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "deleted successful"})
		return
	}
	rules.UpdatedAt = &now
	updateRules, err := h.Repos.UpdateRules(&rules)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't update rules"})
		return
	}
	c.JSON(http.StatusOK, updateRules)
}

func (h *Handler) GetAllProfiles(c *gin.Context) {
	profileName := c.Query("name")
	profiles, err := h.Repos.GetProfiles(profileName)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": ""})
	}
	c.JSON(http.StatusOK, profiles)
}

func (h *Handler) GetProfileRules(c *gin.Context) {
	strProfId := c.Param("id")
	profId, err := strconv.ParseInt(strProfId, 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "can't parse 2 json. Error:" + err.Error()})
		return
	}
	rules, err := h.Repos.GetRules(profId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't get rules. Error:" + err.Error()})
		return
	}
	c.JSON(http.StatusOK, rules)
}
