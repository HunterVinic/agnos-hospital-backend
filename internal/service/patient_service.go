package service

import (
	"agnos-hospital/internal/model"
	"agnos-hospital/internal/repository"
)

type PatientService struct {
	repo       *repository.PatientRepository
	hisService *HISService
}

func NewPatientService(repo *repository.PatientRepository) *PatientService {
	return &PatientService{
		repo:       repo,
		hisService: NewHISService(),
	}
}

// helper to convert string → *string
func stringPtr(s string) *string {
	return &s
}

func (s *PatientService) Search(hospitalID int, nationalID string, passportID string) ([]model.Patient, error) {

	patients, err := s.repo.Search(hospitalID, nationalID, passportID)
	if err != nil {
		return nil, err
	}

	// If found locally
	if len(patients) > 0 {
		return patients, nil
	}

	// If not found locally → call HIS
	if nationalID != "" {
		hisPatient, err := s.hisService.FetchPatient(nationalID)
		if err != nil {
			return patients, nil
		}

		newPatient := model.Patient{
			HospitalID:  hospitalID,
			FirstNameEN: stringPtr(hisPatient.FirstNameEN),
			LastNameEN:  stringPtr(hisPatient.LastNameEN),
			NationalID:  stringPtr(hisPatient.NationalID),
			PassportID:  stringPtr(hisPatient.PassportID),
			Gender:      stringPtr(hisPatient.Gender),
		}

		err = s.repo.Create(&newPatient)
		if err != nil {
			return nil, err
		}

		return []model.Patient{newPatient}, nil
	}

	return patients, nil
}