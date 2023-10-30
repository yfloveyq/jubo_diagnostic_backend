package repository

import (
	"jubo.com/eric/diagnostic/domain"
)

type JuboRepository interface {
	ListPatients() []domain.Patient
	UpdatePatient(patient domain.Patient) (success bool, err error)
	InsertOrder(order domain.Order) (success bool, err error)
	UpdateOrder(order domain.Order) (success bool, err error)
	DeleteOrder(id string) (success bool, err error)
}
