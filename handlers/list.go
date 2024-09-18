package handlers

import (
	"encoding/json"
	"github.com/JohnathanBarb/brimos/models"
	"log"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {

	todos, err := models.GetAll()
	if err != nil {
		log.Printf("Error getting todos: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
