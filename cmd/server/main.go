package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/generalsanjeet/go-model-serve/pkg/db"
)

func main() {
	err := db.InitPostgres()
	if err != nil {
		log.Fatalf("Failed to initialize PostgreSQL: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Go Model Server is running 🎯")
	})

	fmt.Println("🚀 Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

