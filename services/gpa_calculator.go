package services

import "github.com/yaya2127/aastu-academic-portal/models"

// CalculateCumulativeGPA calculates exact grade point average for AASTU transcript
func CalculateCumulativeGPA(grades []models.Grade) float64 {
	if len(grades) == 0 {
		return 0.0
	}

	var totalPoints float64
	var totalECTS float64

	for _, g := range grades {
		totalPoints += g.GradePoints * float64(g.ECTS)
		totalECTS += float64(g.ECTS)
	}

	if totalECTS == 0 {
		return 0.0
	}

	return totalPoints / totalECTS
}
