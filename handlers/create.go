package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/devbersi/finance-go/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var expense models.Expense

	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(expense)

	var resp map[string]any
	if err != nil {
		resp = map[string]any{
			"Error": true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	}
	resp = map[string]any {
		"Error": false,
		"Message": fmt.Sprintf("Expense inserido com sucesso! ID: %v", id),
	}
	
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
} 