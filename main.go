package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yaya2127/aastu-academic-portal/handlers"
)

func main() {
	port := ":8080"

	// Register Go HTTP REST API Routes
	http.HandleFunc("/api/v1/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status":"HEALTHY","system":"AASTU Academic Portal Go Backend","version":"2.0.0"}`)
	})

	http.HandleFunc("/api/v1/student/profile", handlers.GetStudentProfileHandler)

	// Serve Static Single Page Web App
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	fmt.Printf("==================================================\n")
	fmt.Printf("🚀 AASTU Academic Portal Go REST API Server Running\n")
	fmt.Printf("Listening on http://localhost%s\n", port)
	fmt.Printf("==================================================\n")

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
