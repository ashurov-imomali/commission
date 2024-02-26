package router

import (
	"github.com/gin-gonic/gin"
	"main/handler"
)

func GetRouter(h *handler.Handler) *gin.Engine {
	router := gin.New()
	router.GET("/test", h.Test)
	router.POST("/commission-profile", h.CreateCommissionProfile)
	router.PUT("/commission-profile", h.UpdateCommissionProfiles)
	router.PUT("/commission-rules", h.UpdateCommissionRules)
	router.GET("/commission-profile", h.GetAllProfiles)
	router.GET("/commission-profile/:id", h.GetProfileRules)
	return router
}
