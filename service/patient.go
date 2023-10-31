package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"jubo.com/eric/diagnostic/domain"
	"jubo.com/eric/diagnostic/repository"
)

type PatientService struct {
	juboRepository repository.JuboRepository
}

func NewPatientService() *PatientService {
	return &PatientService{repository.NewMongoJuboRepository()}
}

func (service *PatientService) ListPatients() []domain.Patient {
	return service.juboRepository.ListPatients()
}

func (service *PatientService) UpdatePatient(patient domain.Patient) (success bool, err error) {
	return service.juboRepository.UpdatePatient(patient)
}

func (service *PatientService) InsertOrder(patient domain.Patient, order domain.Order) (o domain.Order, err error) {
	id, err := service.juboRepository.InsertOrder(order)
	if err == nil {
		order.ID = id
		patient.OrderList = append(patient.OrderList, id)
		success, err := service.juboRepository.UpdatePatient(patient)
		if success {
			return order, err
		}
	}
	return domain.Order{}, err
}

func (service *PatientService) UpdateOrder(order domain.Order) (success bool, err error) {
	return service.juboRepository.UpdateOrder(order)
}

func (service *PatientService) DeleteOrder(patient domain.Patient, id primitive.ObjectID) (success bool, err error) {
	s, err := service.juboRepository.DeleteOrder(id)
	if s {
		length := len(patient.OrderList)
		var newList []primitive.ObjectID
		for i := 0; i < length; i++ {
			existing := patient.OrderList[i]
			if id != patient.OrderList[i] {
				newList = append(newList, existing)
			}
		}
		patient.OrderList = newList
		return service.juboRepository.UpdatePatient(patient)
	}
	return false, err
}
