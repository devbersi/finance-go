package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/devbersi/finance-go/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	expenses, err := models.GetAll()
	if err != nil {
		log.Printf("Erro ao listar registros")
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expenses)
}