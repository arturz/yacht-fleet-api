package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"rest/models"
	"strconv"
)

func ListMarinas(w http.ResponseWriter, r *http.Request) {
	marinas := models.GetMarinas()
	json.NewEncoder(w).Encode(marinas)
}

func GetMarina(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	marina, ok := models.Marinas[id]

	if !ok {
		http.Error(w, "Marina not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(marina)
}

func CreateMarina(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newMarina := models.NewMarina(input.Name)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", r.URL.Path+"/"+strconv.Itoa(newMarina.ID))
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMarina)
}

func UpdateMarina(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	marina, ok := models.Marinas[id]

	if !ok {
		http.Error(w, "Marina not found", http.StatusNotFound)
		return
	}

	var input struct {
		Name string `json:"name"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	marina.Name = input.Name
	models.Marinas[id] = marina

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(marina)
}

func DeleteMarina(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	_, ok := models.Marinas[id]

	if !ok {
		http.Error(w, "Marina not found", http.StatusNotFound)
		return
	}

	delete(models.Marinas, id)

	w.WriteHeader(http.StatusNoContent)
}

func ListYachtsInMarina(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	yachts := []models.Yacht{}
	for _, yacht := range models.Yachts {
		if yacht.MarinaID == id {
			yachts = append(yachts, yacht)
		}
	}
	json.NewEncoder(w).Encode(yachts)
}
