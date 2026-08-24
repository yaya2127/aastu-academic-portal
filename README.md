# 🎓 AASTU Student & Academic Management Portal (Go Microservices & React)

![Go Language](https://img.shields.io/badge/Go-Golang%201.22-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![TypeScript](https://img.shields.io/badge/TypeScript-React%2018-3178C6?style=for-the-badge&logo=typescript&logoColor=white)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Database-4169E1?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Containerized-2496ED?style=for-the-badge&logo=docker&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green?style=for-the-badge)

A high-performance enterprise academic microservice web application engineered for **Addis Ababa Science and Technology University (AASTU)** students and faculty. Features a **Go (Golang)** REST API backend, **React + TypeScript** frontend, and containerized **PostgreSQL** database service.

---

## ✨ Enterprise Architecture & Capabilities

- ⚡ **High-Performance Go (Golang) REST API**: Lightweight, concurrent HTTP server routes (`/api/v1/student/profile`, `/api/v1/health`) handling student profile data, course registration validation, and GPA calculation logic.
- 🔷 **React 18 + TypeScript Frontend Architecture**: Strictly typed component architecture (`App.tsx`, `StudentProfile`, `Course`, `GradeRecord`) with custom state hooks and API integration.
- 🐘 **PostgreSQL Relational Persistence**: Relational database schema for student profiles, course catalogs, ECTS credit logs, and letter grade records.
- 🐳 **Docker & Docker Compose Containerization**: Multi-stage Dockerfile and `docker-compose.yml` defining the Go microservice API container and isolated PostgreSQL instance.

---

## 🛠️ Advanced Tech Stack

| Component | Technology |
| :--- | :--- |
| **Backend API** | Go (Golang 1.22), Standard `net/http` Engine, RESTful Architecture |
| **Frontend Web** | React 18, TypeScript, HTML5, CSS3 |
| **Database** | PostgreSQL 16 Relational Engine |
| **Containerization** | Docker, Docker Compose, Multi-stage Alpine builds |

---

## 📁 Repository Structure

```text
aastu-academic-portal/
├── main.go                     # Go HTTP Server Entry Point & Router
├── go.mod                      # Go Module Definition
├── handlers/
│   └── student.go              # Go HTTP Handler Controllers
├── models/
│   └── student.go              # Go Struct Data Models
├── src/
│   ├── App.tsx                 # React + TypeScript App Component
│   └── types/
│       └── index.ts            # TypeScript Interfaces
├── Dockerfile                  # Multi-Stage Docker Container Build
├── docker-compose.yml          # Container Orchestration (Go + PostgreSQL)
└── README.md                   # Repository Documentation
```

---

## 🚀 Quick Start & Local Execution

### Option A: Run Go Backend Server Locally
```bash
# Clone the repository
git clone https://github.com/yaya2127/aastu-academic-portal.git
cd aastu-academic-portal

# Run Go REST API server
go run main.go
```
Open `http://localhost:8080` in your web browser.

### Option B: Launch with Docker Compose
```bash
docker-compose up --build
```

---

## 👤 Author

Developed by **Yared Kinetibeb Tesfaye**
- **Role**: 5th-Year Computer Engineering Senior Student at AASTU
- **GitHub**: [@yaya2127](https://github.com/yaya2127)
- **LinkedIn**: [Yared Kinetibeb](https://www.linkedin.com/in/yared-kinetibeb-3b788b350/)
- **Email**: kinetibebyared@gmail.com


## Go REST API Endpoints
- GET /api/v1/health -> Health Check
- GET /api/v1/student/profile -> Student Record


## Transcript PDF Export
- Automated official academic transcript builder


## GPA Calculation Engine
- Weighted grade point average algorithm


## JWT Authentication Middleware
- Secure token validation


## REST API Payload Schemas
- Student profile JSON structure

<!-- Go build cache config -->

<!-- Routing performance tuning -->
