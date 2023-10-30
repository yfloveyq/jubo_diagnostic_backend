package service

import (
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

func (service *PatientService) InsertOrder(patient domain.Patient, order domain.Order) (success bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (service *PatientService) UpdateOrder(order domain.Order) (success bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (service *PatientService) DeleteOrder(patient domain.Patient, id string) (successful bool, err error) {
	//TODO implement me
	panic("implement me")
}
