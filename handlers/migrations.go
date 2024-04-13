package handlers

import (
	"github.com/gorilla/mux"
	//"github.com/stretchr/testify/assert"
	"encoding/json"
	"log"
	"net/http"
	"rest/models"
	"strconv"
)

func ListMigrations(w http.ResponseWriter, r *http.Request) {
	migrations := models.GetMigrations()
	json.NewEncoder(w).Encode(migrations)
}

func GetMigration(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	migration, ok := models.Migrations[id]

	if !ok {
		http.Error(w, "Migration not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(migration)
}

func CreateMigration(w http.ResponseWriter, r *http.Request) {
	tokenStr := r.URL.Query().Get("token")

	token, err := strconv.Atoi(tokenStr)

	if err != nil {
		// log error to console
		log.Println(err)

		http.Error(w, "Invalid token", http.StatusBadRequest)
		return
	}

	if models.IsTokenUsed(token) {
		http.Error(w, "Used token", http.StatusBadRequest)
		return
	}

	var input struct {
		YachtID  int `json:"yacht_id"`
		MarinaID int `json:"marina_id"`
	}

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	migration, err := models.NewMigration(input.YachtID, input.MarinaID)

	models.UseToken(token)

	if err != nil {
		http.Error(w, "Cannot create migration", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", r.URL.Path+"/"+strconv.Itoa(migration.ID))
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(migration)
}
