# 🏥 Agnos Hospital Backend Service

Production-ready backend service built with Go (Gin), PostgreSQL, Docker, and Nginx.

Implements:

- 🔐 Hospital-scoped staff authentication (JWT-based)
- 👥 Multi-tenant patient search
- 🔎 External HIS fallback integration
- 🐳 Containerized deployment with reverse proxy
- 🧪 Pre-seeded multi-hospital demo data

---

## 🚀 Tech Stack

| Layer                        | Technology              |
|------------------------------|-------------------------|
| Language                     | Go 1.25                 | 
| Framework                    | Gin                     |
| Database                     | PostgreSQL 15           |
| Authentication               | JWT (HS256)             |
| Password Hashing             | bcrypt                  |
| Reverse Proxy                | Nginx                   |
| Containerization             | Docker & Docker Compose |

---

## 🏗 System Architecture

Client  
↓  
Nginx (Port 80)  
↓  
Gin Application (Port 8080)  
↓  
PostgreSQL  
↓  
External HIS API  

---

## 🧠 Architecture Highlights

### Clean Layered Design

Handler → Service → Repository → Database

Benefits:

- Clear separation of concerns
- Easier testing
- Business logic isolation
- Scalable structure

---

### Multi-Tenant (Hospital-Level Isolation)

JWT payload contains:

{
  "staff_id": 1,
  "hospital_id": 2,
  "exp": 1700000000
}

Every patient query enforces:

WHERE hospital_id = ?

This guarantees strict hospital data isolation.

Staff from Hospital 1 cannot see patients from Hospital 2.

---

## 🧪 Pre-Seeded Demo Data

The database automatically seeds:

### 🏥 Hospitals

1. Bangkok General Hospital
2. Chiang Mai Medical Center
3. Phuket International Hospital

---

### 👨‍⚕️ Staff Accounts

Password for ALL accounts: 1234

| Username        | Hospital |
|-----------------|----------|
| admin_bkk      | Hospital 1 |
| doctor_cm      | Hospital 2 |
| nurse_phuket   | Hospital 3 |

---

### 👥 Sample Patients

Each hospital has 2 pre-seeded patients.

You can test multi-tenant isolation immediately.

---

## 📦 Project Structure

agnos-hospital/

├── cmd/
│   └── main.go
│
├── internal/
│   ├── config/
│   ├── model/
│   ├── repository/
│   ├── service/
│   ├── handler/
│   └── middleware/
│
├── migrations/
│   └── init.sql
│
├── nginx/
│   └── nginx.conf
│
├── Dockerfile
├── docker-compose.yml
└── README.md

---

## 🐳 Running the Project

IMPORTANT: Because seed data runs only on first DB creation:

docker compose down -v
docker compose up --build

Application available at:

http://localhost

---

## 🔐 Authentication Flow

1. Login using seeded account
2. Receive JWT token
3. Access protected routes with:

Authorization: Bearer <token>

---

## 📡 API Specification

### Health Check

GET /health

---

### Staff Login

POST /staff/login

Example:

curl -X POST http://localhost/staff/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin_bkk","password":"1234"}'

---

### Search Patient (Protected)

GET /patient/search

Headers:

Authorization: Bearer <token>

Query options:

?national_id=1111111111111  
?passport_id=AA111111  

---

## 🧪 Multi-Tenant Demo Example

Login as Bangkok staff:

admin_bkk

Search patient with national_id:

1111111111111  ✅ (visible)

Try searching patient from Hospital 2:

2222222222221  ❌ (not visible)

This proves strict hospital-level data isolation.

---

## 🗄 Database Design

- Foreign key enforcement
- Cascade delete protection
- Unique national_id and passport_id
- Indexed search fields
- Seeded multi-hospital environment

---

## 🛡 Security Features

- bcrypt password hashing
- JWT expiration (24h)
- Middleware-based authorization
- Parameterized SQL queries
- Reverse proxy separation
- Stateless authentication
- Multi-tenant isolation

---

## 📈 Scalability Design

- Stateless JWT (horizontal scaling ready)
- Nginx reverse proxy
- Repository abstraction
- HIS integration layer
- Dockerized deployment

---

## ⚖ Tradeoffs

- No caching layer (Redis optional)
- No migration tool (Goose recommended)
- No rate limiting (can be added in Nginx)
- HIS simplified for assignment scope

---

## 👨‍💻 Author

Sheshehang Limbu  
Backend Developer (Golang)

---

## 🏆 What This Project Demonstrates

- Clean architecture understanding
- Secure multi-tenant backend design
- Production-ready containerization
- Relational database integrity
- External system integration pattern
- Professional-level backend structuring

