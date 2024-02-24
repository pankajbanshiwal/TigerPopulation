package App

import (
	sightingController "TigerPopulation/Controllers/Sightings"
	tigersController "TigerPopulation/Controllers/Tigers"
	controller "TigerPopulation/Controllers/Users"
)

func MapDevUrls() {
	var dev = router.Group("/dev")
	var devApp = router.Group(dev.BasePath() + "/tigerMonitor")
	devApp.PUT("/user/create", controller.CreateUser)
	devApp.POST("/user/login", controller.Login)
	devApp.PUT("/tiger/create", tigersController.CreateTiger)
	devApp.GET("/tigers", tigersController.GetAllTigers)
	devApp.POST("/create/sighting", sightingController.CreateSighting)
	devApp.GET("/tiger/sightings", sightingController.GetTigerSightings)
}
