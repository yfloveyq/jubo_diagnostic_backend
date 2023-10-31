package main

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	restfulController.Bind(controller.PUT, "/patients/:id", func(c *gin.Context) {
		id, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid patient ID."})
			return
		}
		patient := domain.Patient{}
		err = c.ShouldBind(&patient)
		if err == nil && patient.ID == id {
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

	restfulController.Bind(controller.POST, "/orders", func(c *gin.Context) {
		orderRequest := domain.OrderRequest{}
		err := c.ShouldBind(&orderRequest)
		if err == nil {
			patient := orderRequest.Patient
			order := orderRequest.Order
			o, err := patientService.InsertOrder(patient, order)
			if err == nil {
				c.JSON(200, o)
			} else {
				c.JSON(500, gin.H{"error": err.Error()})
			}
		} else {
			c.JSON(400, gin.H{"error": "Invalid request data."})
		}
	})

	restfulController.Bind(controller.PUT, "/orders/:id", func(c *gin.Context) {
		id, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid order ID."})
			return
		}
		order := domain.Order{}
		err = c.ShouldBind(&order)
		if err == nil && order.ID == id {
			o, err := patientService.UpdateOrder(order)
			if o {
				c.JSON(200, order)
			} else {
				c.JSON(500, gin.H{"error": err.Error()})
			}
		} else {
			c.JSON(400, gin.H{"error": "Invalid order data."})
		}
	})

	restfulController.Bind(controller.DELETE, "/orders/:id", func(c *gin.Context) {
		id, err := primitive.ObjectIDFromHex(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid order ID."})
			return
		}
		patient := domain.Patient{}
		err = c.ShouldBind(&patient)
		if err == nil {
			o, err := patientService.DeleteOrder(patient, id)
			if o {
				c.JSON(200, patient)
			} else {
				c.JSON(500, gin.H{"error": err.Error()})
			}
		} else {
			c.JSON(400, gin.H{"error": "Invalid request data."})
		}
	})

	restfulController.Start()
}
