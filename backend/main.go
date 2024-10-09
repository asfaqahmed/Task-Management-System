// backend/main.go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"task-management-app/handlers"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
)

var db *sql.DB

func main() {
	// Initialize MSSQL database connection
	connStr := "server=localhost;user id=sa;password=yourpassword;database=TaskDB"
	var err error
	db, err = sql.Open("sqlserver", connStr)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	defer db.Close()

	// Initialize Router
	router := mux.NewRouter()
	router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	// Start server
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
