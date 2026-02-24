package model

import "time"

type Patient struct {
	ID          int        `json:"id"`
	HospitalID  int        `json:"hospital_id"`
	FirstNameEN *string    `json:"first_name_en"`
	LastNameEN  *string    `json:"last_name_en"`
	NationalID  *string    `json:"national_id"`
	PassportID  *string    `json:"passport_id"`
	PhoneNumber *string    `json:"phone_number"`
	Email       *string    `json:"email"`
	Gender      *string    `json:"gender"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	CreatedAt   time.Time  `json:"created_at"`
}