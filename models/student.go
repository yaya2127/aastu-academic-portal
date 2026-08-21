package models

import "time"

// StudentProfile represents a university student record at AASTU
type StudentProfile struct {
	StudentID   string             `json:"student_id"`
	FullName    string             `json:"full_name"`
	Department  string             `json:"department"`
	YearLevel   int                `json:"year_level"`
	CGPA        float64            `json:"cgpa"`
	Enrollments []CourseEnrollment `json:"enrollments"`
	CreatedAt   time.Time          `json:"created_at"`
}

// CourseEnrollment represents a registered course unit
type CourseEnrollment struct {
	CourseCode  string  `json:"course_code"`
	Title       string  `json:"title"`
	CreditHours int     `json:"credit_hours"`
	Grade       string  `json:"grade"`
	GradePoint  float64 `json:"grade_point"`
}

// CalculateGPA calculates the weighted GPA from course grade points
func CalculateGPA(courses []CourseEnrollment) float64 {
	if len(courses) == 0 {
		return 0.0
	}
	var totalPoints float64
	var totalCredits int

	for _, c := range courses {
		totalPoints += c.GradePoint * float64(c.CreditHours)
		totalCredits += c.CreditHours
	}

	if totalCredits == 0 {
		return 0.0
	}
	return totalPoints / float64(totalCredits)
}
