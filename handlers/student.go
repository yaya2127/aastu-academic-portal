package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/yaya2127/aastu-academic-portal/models"
)

// GetStudentProfileHandler returns Yared Kinetibeb's AASTU academic profile in JSON
func GetStudentProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	student := models.Student{
		ID:            "ETS0001/15",
		FullName:      "Yared Kinetibeb Tesfaye",
		Department:    "Computer Engineering",
		YearStanding:  5,
		CumulativeGPA: 3.78,
		CompletedECTS: 142,
		ActiveCourses: []models.Course{
			{Code: "CoEng-5101", Title: "Embedded Systems Architecture & Design", ECTS: 6, Instructor: "Dr. Abebe T.", Status: "Registered"},
			{Code: "CoEng-5102", Title: "Advanced Distributed Systems & Web APIs", ECTS: 5, Instructor: "Eng. Samuel M.", Status: "Registered"},
			{Code: "CoEng-5103", Title: "Agentic AI & Machine Learning Systems", ECTS: 6, Instructor: "Dr. Helen G.", Status: "Registered"},
			{Code: "CoEng-5104", Title: "Senior Engineering Capstone Project I", ECTS: 8, Instructor: "Department Committee", Status: "Registered"},
		},
		GradesHistory: []models.Grade{
			{CourseCode: "CoEng-4101", CourseTitle: "Software Engineering & Architecture", ECTS: 6, GradeLetter: "A+", GradePoints: 4.0, Semester: "Year 4 Sem II"},
			{CourseCode: "CoEng-4102", CourseTitle: "Microcontrollers & Hardware Interfacing", ECTS: 6, GradeLetter: "A", GradePoints: 4.0, Semester: "Year 4 Sem II"},
			{CourseCode: "CoEng-4103", CourseTitle: "Database Management Systems", ECTS: 5, GradeLetter: "A", GradePoints: 4.0, Semester: "Year 4 Sem I"},
		},
	}

	json.NewEncoder(w).Encode(student)
}
