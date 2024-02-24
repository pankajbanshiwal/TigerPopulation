package App

import (
	"TigerPopulation/Utils/dbConfig"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupNoRoute() {
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
}

func healthCheck() {
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})
}

func readyCheck() {
	var err error
	router.GET("/readyz", func(c *gin.Context) {
		if dbConfig.DB == nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status": "NIL",
			})
		} else if err = dbConfig.DB.Ping(); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"status": "PING FAILED",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "READY",
			})
		}
	})
}
func SetupRouteUtils() {
	setupNoRoute()
	healthCheck()
	readyCheck()
}
