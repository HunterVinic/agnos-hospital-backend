package model

import "time"

type Staff struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"-"`
	HospitalID   int       `json:"hospital_id"`
	CreatedAt    time.Time `json:"created_at"`
}