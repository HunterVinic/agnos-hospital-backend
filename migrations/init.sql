-- =====================================
-- TABLES
-- =====================================

CREATE TABLE IF NOT EXISTS hospitals (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS staff (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    hospital_id INT REFERENCES hospitals(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS patients (
    id SERIAL PRIMARY KEY,
    hospital_id INT REFERENCES hospitals(id) ON DELETE CASCADE,
    first_name_en VARCHAR(100),
    last_name_en VARCHAR(100),
    national_id VARCHAR(20) UNIQUE,
    passport_id VARCHAR(20) UNIQUE,
    phone_number VARCHAR(20),
    email VARCHAR(255),
    gender VARCHAR(1),
    date_of_birth DATE,
    created_at TIMESTAMP DEFAULT NOW()
);

-- =====================================
-- INDEXES
-- =====================================

CREATE INDEX IF NOT EXISTS idx_patients_national_id ON patients(national_id);
CREATE INDEX IF NOT EXISTS idx_patients_passport_id ON patients(passport_id);
CREATE INDEX IF NOT EXISTS idx_patients_hospital_id ON patients(hospital_id);
CREATE INDEX IF NOT EXISTS idx_staff_username ON staff(username);

-- =====================================
-- SEED HOSPITALS
-- =====================================

INSERT INTO hospitals (id, name) VALUES
(1, 'Bangkok General Hospital'),
(2, 'Chiang Mai Medical Center'),
(3, 'Phuket International Hospital')
ON CONFLICT (id) DO NOTHING;

-- =====================================
-- SEED STAFF
-- Password for all accounts: 1234
-- bcrypt hash generated using Go bcrypt
-- =====================================

INSERT INTO staff (username, password_hash, hospital_id) VALUES
(
    'admin_bkk',
    '$2a$10$7EqJtq98hPqEX7fNZaFWoOaQpV8V6E9E2lY7W7Zl5VwH1Q6QYF7eK',
    1
),
(
    'doctor_cm',
    '$2a$10$7EqJtq98hPqEX7fNZaFWoOaQpV8V6E9E2lY7W7Zl5VwH1Q6QYF7eK',
    2
),
(
    'nurse_phuket',
    '$2a$10$7EqJtq98hPqEX7fNZaFWoOaQpV8V6E9E2lY7W7Zl5VwH1Q6QYF7eK',
    3
)
ON CONFLICT (username) DO NOTHING;

-- =====================================
-- SEED PATIENTS
-- =====================================

INSERT INTO patients (
    hospital_id,
    first_name_en,
    last_name_en,
    national_id,
    passport_id,
    phone_number,
    email,
    gender,
    date_of_birth
) VALUES

-- Hospital 1 Patients
(
    1,
    'John',
    'Doe',
    '1111111111111',
    'AA111111',
    '0811111111',
    'john.doe@email.com',
    'M',
    '1990-01-01'
),
(
    1,
    'Sarah',
    'Connor',
    '1111111111112',
    'AA111112',
    '0811111112',
    'sarah.connor@email.com',
    'F',
    '1988-04-12'
),

-- Hospital 2 Patients
(
    2,
    'Michael',
    'Tan',
    '2222222222221',
    'BB222221',
    '0822222221',
    'michael.tan@email.com',
    'M',
    '1995-03-20'
),
(
    2,
    'Anna',
    'Lee',
    '2222222222222',
    'BB222222',
    '0822222222',
    'anna.lee@email.com',
    'F',
    '1993-07-15'
),

-- Hospital 3 Patients
(
    3,
    'David',
    'Wilson',
    '3333333333331',
    'CC333331',
    '0833333331',
    'david.w@email.com',
    'M',
    '1985-09-09'
),
(
    3,
    'Emily',
    'Clark',
    '3333333333332',
    'CC333332',
    '0833333332',
    'emily.c@email.com',
    'F',
    '1991-11-11'
)

ON CONFLICT (national_id) DO NOTHING;