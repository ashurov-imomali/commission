package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/db"
	"main/models"
	"net/http"
)

type Handler struct {
	Repos *db.Repository
}

func GetHandler(repository *db.Repository) *Handler {
	return &Handler{repository}
}

func (h *Handler) Test(c *gin.Context) {
	c.String(200, "oook")
}

func (h *Handler) CreateCommissionProfile(c *gin.Context) {
	var request models.ProfileCreatRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "can't parse 2 json"})
		return
	}
	var userId int64 = 1
	request.Profile.CreatedBy = userId
	profileId, err := h.Repos.CreateProfile(&request.Profile)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't create profile"})
		return
	}
	err = h.Repos.CreateRules(request.Rules, profileId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't add new rules"})
		return
	}
	c.Status(201)
}

func (h *Handler) UpdateCommissionProfiles(c *gin.Context) {
	var updProfile models.CommissionProfiles
	err := c.ShouldBindJSON(&updProfile)
	if err != nil {
		log.Println(err)
	}

}
