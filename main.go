package main

import (
	"github.com/gin-gonic/gin"
	"jubo.com/eric/diagnostic/controller"
	"jubo.com/eric/diagnostic/service"
)

func main() {
	restfulController := controller.NewRestfulController()

	patientService := service.NewPatientService()

	restfulController.Bind(controller.GET, "/patients", func(c *gin.Context) {
		c.JSON(200, patientService.ListPatients())
	})

	restfulController.Start()
}
