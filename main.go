package main

import (
	"github.com/gin-gonic/gin"
	"jubo.com/eric/diagnostic/controller"
	"jubo.com/eric/diagnostic/domain"
	"jubo.com/eric/diagnostic/service"
)

func main() {
	restfulController := controller.NewRestfulController()

	patientService := service.NewPatientService()

	restfulController.Bind(controller.GET, "/patients", func(c *gin.Context) {
		c.JSON(200, patientService.ListPatients())
	})

	restfulController.Bind(controller.PUT, "/patients", func(c *gin.Context) {
		patient := domain.Patient{}
		err := c.ShouldBind(&patient)
		if err == nil {
			success, err := patientService.UpdatePatient(patient)
			if success {
				c.JSON(200, patient)
			} else {
				c.JSON(500, gin.H{"error": err.Error()})
			}
		} else {
			c.JSON(400, gin.H{"error": "Invalid patient data."})
		}

	})

	restfulController.Start()
}
