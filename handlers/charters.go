package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"rest/models"
	"strconv"
)

func ListCharters(w http.ResponseWriter, r *http.Request) {
	charters := models.GetCharters()
	json.NewEncoder(w).Encode(charters)
}

func GetCharter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	charter, ok := models.Charters[id]

	if !ok {
		http.Error(w, "Charter not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(charter)
}

func CreateCharter(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Captain string `json:"captain"`
		YachtID int    `json:"yacht_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newCharter := models.NewCharter(input.Captain, input.YachtID)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", r.URL.Path+"/"+strconv.Itoa(newCharter.ID))
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCharter)
}

func UpdateCharter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	charter, ok := models.Charters[id]

	if !ok {
		http.Error(w, "Charter not found", http.StatusNotFound)
		return
	}

	var input struct {
		Captain string `json:"captain"`
		YachtID int    `json:"yacht_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	charter.Captain = input.Captain
	charter.YachtID = input.YachtID
	models.Charters[id] = charter

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(charter)
}

func DeleteCharter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	_, ok := models.Charters[id]

	if !ok {
		http.Error(w, "Charter not found", http.StatusNotFound)
		return
	}

	delete(models.Charters, id)

	w.WriteHeader(http.StatusNoContent)
}
