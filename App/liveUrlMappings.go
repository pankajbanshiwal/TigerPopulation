package App

import (
	sightingController "TigerPopulation/Controllers/Sightings"
	tigersController "TigerPopulation/Controllers/Tigers"
	controller "TigerPopulation/Controllers/Users"
)

func MapLiveUrls() {
	var v1 = router.Group("/v1")
	var liveApp = router.Group(v1.BasePath() + "/tigerMonitor")
	liveApp.PUT("/user/create", controller.CreateUser)
	liveApp.POST("/user/login", controller.Login)
	liveApp.PUT("/tiger/create", tigersController.CreateTiger)
	liveApp.GET("/tigers", tigersController.GetAllTigers)
	liveApp.POST("/create/sighting", sightingController.CreateSighting)
	liveApp.GET("/tiger/sightings", sightingController.GetTigerSightings)
}
