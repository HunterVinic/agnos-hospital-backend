package repository

import (
	"agnos-hospital/internal/model"
	"database/sql"
	"fmt"
)

type PatientRepository struct {
	DB *sql.DB
}

func NewPatientRepository(db *sql.DB) *PatientRepository {
	return &PatientRepository{DB: db}
}

func (r *PatientRepository) Search(hospitalID int, nationalID string, passportID string) ([]model.Patient, error) {

	query := `SELECT 
		id, hospital_id, first_name_en, last_name_en,
		national_id, passport_id, phone_number, email,
		gender, date_of_birth, created_at
		FROM patients
		WHERE hospital_id = $1`

	args := []interface{}{hospitalID}
	argIndex := 2

	if nationalID != "" {
		query += fmt.Sprintf(" AND national_id = $%d", argIndex)
		args = append(args, nationalID)
		argIndex++
	}

	if passportID != "" {
		query += fmt.Sprintf(" AND passport_id = $%d", argIndex)
		args = append(args, passportID)
		argIndex++
	}

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []model.Patient

	for rows.Next() {
		var p model.Patient

		var firstName, lastName, natID, passID, phone, email, gender sql.NullString
		var dob sql.NullTime

		err := rows.Scan(
			&p.ID,
			&p.HospitalID,
			&firstName,
			&lastName,
			&natID,
			&passID,
			&phone,
			&email,
			&gender,
			&dob,
			&p.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		if firstName.Valid {
			p.FirstNameEN = &firstName.String
		}
		if lastName.Valid {
			p.LastNameEN = &lastName.String
		}
		if natID.Valid {
			p.NationalID = &natID.String
		}
		if passID.Valid {
			p.PassportID = &passID.String
		}
		if phone.Valid {
			p.PhoneNumber = &phone.String
		}
		if email.Valid {
			p.Email = &email.String
		}
		if gender.Valid {
			p.Gender = &gender.String
		}
		if dob.Valid {
			p.DateOfBirth = &dob.Time
		}

		patients = append(patients, p)
	}

	return patients, nil
}


func (r *PatientRepository) Create(patient *model.Patient) error {
	query := `
		INSERT INTO patients
		(hospital_id, first_name_en, last_name_en, national_id, passport_id, gender)
		VALUES ($1,$2,$3,$4,$5,$6)
		RETURNING id, created_at
	`

	return r.DB.QueryRow(
		query,
		patient.HospitalID,
		patient.FirstNameEN,
		patient.LastNameEN,
		patient.NationalID,
		patient.PassportID,
		patient.Gender,
	).Scan(&patient.ID, &patient.CreatedAt)
}