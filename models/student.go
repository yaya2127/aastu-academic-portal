package models

// Student represents an AASTU Computer Engineering undergraduate student profile
type Student struct {
	ID             string    `json:"id"`
	FullName       string    `json:"full_name"`
	Department     string    `json:"department"`
	YearStanding   int       `json:"year_standing"`
	CumulativeGPA  float64   `json:"cumulative_gpa"`
	CompletedECTS  int       `json:"completed_ects"`
	ActiveCourses  []Course  `json:"active_courses"`
	GradesHistory  []Grade   `json:"grades_history"`
}

// Course represents an academic course module
type Course struct {
	Code       string `json:"code"`
	Title      string `json:"title"`
	ECTS       int    `json:"ects"`
	Instructor string `json:"instructor"`
	Status     string `json:"status"`
}

// Grade represents a completed course grade record
type Grade struct {
	CourseCode  string  `json:"course_code"`
	CourseTitle string  `json:"course_title"`
	ECTS        int     `json:"ects"`
	GradeLetter string  `json:"grade_letter"`
	GradePoints float64 `json:"grade_points"`
	Semester    string  `json:"semester"`
}
