package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tomasarras/GoLang/controller/agencyController"
)

func Start() {
	r := gin.Default()

	r.POST("/agencies", agencyController.SaveAgencyHandler)
	r.GET("/agencies/:id", agencyController.FindByIdAgencyHandler)
	r.GET("/agencies", agencyController.FindAllAgencyHandler)
	r.PUT("/agencies/:id", agencyController.UpdateAgencyHandler)
	r.DELETE("/agencies/:id", agencyController.RemoveAgencyHandler)

	/*
		r.POST("/agencies/:idAgency/flights", flightController.SaveFlightHandler)
		r.GET("/agencies/flights/:id", agencyController.FindByIdAgencyHandler)
		r.PUT("/agencies/flights/:id", agencyController.UpdateAgencyHandler)
		r.DELETE("/agencies/flights/:id", agencyController.RemoveAgencyHandler)
		r.GET("/agencies/flights", agencyController.FindAllAgencyHandler)
	*/
	r.Run()
}
