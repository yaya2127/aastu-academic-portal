package services

func ConvertGradeToPoint(grade string) float64 {
	if grade == "A" { return 4.0 }
	return 3.0
}
