package repository

import (
	"agnos-hospital/internal/model"
	"database/sql"
)

type StaffRepository struct {
	DB *sql.DB
}

func NewStaffRepository(db *sql.DB) *StaffRepository {
	return &StaffRepository{DB: db}
}

func (r *StaffRepository) Create(staff *model.Staff) error {
	query := `
		INSERT INTO staff (username, password_hash, hospital_id)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`
	return r.DB.QueryRow(
		query,
		staff.Username,
		staff.PasswordHash,
		staff.HospitalID,
	).Scan(&staff.ID, &staff.CreatedAt)
}

func (r *StaffRepository) FindByUsername(username string) (*model.Staff, error) {
	query := `
		SELECT id, username, password_hash, hospital_id, created_at
		FROM staff
		WHERE username = $1
	`

	row := r.DB.QueryRow(query, username)

	staff := &model.Staff{}
	err := row.Scan(
		&staff.ID,
		&staff.Username,
		&staff.PasswordHash,
		&staff.HospitalID,
		&staff.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return staff, nil
}