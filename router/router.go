package router

import (
	"net/http"

	"github.com/dimfeld/httptreemux/v5"
	"github.com/tomasarras/GoLang/controller/agencyController"
	"github.com/tomasarras/GoLang/controller/flightController"
)

func Start() {
	router := httptreemux.New()

	router.POST("/agencies", agencyController.SaveAgencyHandler)
	router.GET("/agencies/:id", agencyController.FindByIdAgencyHandler)
	router.GET("/agencies", agencyController.FindAllAgencyHandler)
	router.PUT("/agencies/:id", agencyController.UpdateAgencyHandler)
	router.DELETE("/agencies/:id", agencyController.RemoveAgencyHandler)

	router.POST("/agencies/:idAgency/flights", flightController.SaveFlightHandler)
	router.GET("/agencies/flights/:id", flightController.FindByIdFlightHandler)
	router.GET("/agencies/flights", flightController.FindAllFlightHandler)
	router.GET("/agencies/:idAgency/flights", flightController.FindAllFlightByAgencyHandler)
	router.PUT("/agencies/flights/:id", flightController.UpdateFlightHandler)
	router.DELETE("/agencies/flights/:id", flightController.RemoveFlightHandler)

	http.ListenAndServe(":8080", router)
}
