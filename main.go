package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Response struct {
	CurrentTime string `json:"current_time"`
}

func main() {
	// Database connection setup
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_api")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Test the connection
	if err = db.Ping(); err != nil {
		log.Fatalf("Database is not reachable: %v", err)
	}

	http.HandleFunc("/current-time", func(w http.ResponseWriter, r *http.Request) {
		// Get the current time in Toronto
		loc, err := time.LoadLocation("America/Toronto")
		if err != nil {
			http.Error(w, "Error loading timezone", http.StatusInternalServerError)
			log.Println("Timezone error:", err)
			return
		}
		now := time.Now().In(loc)

		// Insert the current time into the database
		_, err = db.Exec("INSERT INTO time_log (timestamp) VALUES (?)", now)
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			log.Println("Database error:", err)
			return
		}

		// Create the JSON response
		response := Response{CurrentTime: now.Format(time.RFC3339)}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// Start the server
	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
