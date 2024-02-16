package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"main.go/models"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var resp map[string]any

	//check if user is deleted
	u, err := models.GetUser(user.Uuid)
	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar buscar o usuário %v", err),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}
	if u.Deleted {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("O usuário não pode ser atualizado pois está deletado %v", err),
		}
		json.NewEncoder(w).Encode(resp)
		return
	}
	_, err = models.UpdateUser(user)
	if err != nil {
		log.Printf("Erro ao fazer update do usuário: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	resp = map[string]any{
		"Error":   false,
		"Message": fmt.Sprintf("Usuário atualizado com sucesso! ID: %s", user.Uuid),
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
