package handlers

import "net/http"

func GetStudentProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"student_id":"ETS-0452/14"}`))
}
