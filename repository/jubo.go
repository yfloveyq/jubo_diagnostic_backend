package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"jubo.com/eric/diagnostic/domain"
)

type JuboRepository interface {
	ListPatients() []domain.Patient
	UpdatePatient(patient domain.Patient) (success bool, err error)
	InsertOrder(order domain.Order) (id primitive.ObjectID, err error)
	UpdateOrder(order domain.Order) (success bool, err error)
	DeleteOrder(id primitive.ObjectID) (success bool, err error)
}
