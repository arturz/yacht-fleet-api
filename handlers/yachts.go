package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"rest/models"
	"rest/utils"
	"strconv"
)

func ListYachts(w http.ResponseWriter, r *http.Request) {
	yachts := models.GetYachts()

	paginatedYachts, err := utils.Paginate(r, yachts)
	if err != nil {
		http.Error(w, "Invalid pagination parameters", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paginatedYachts)
}

func GetYacht(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	yacht, ok := models.Yachts[id]

	if !ok {
		http.Error(w, "Yacht not found", http.StatusNotFound)
		return
	}

	etag := models.HashYacht(id)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Etag", etag)

	json.NewEncoder(w).Encode(yacht)
}

func CreateYacht(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		MarinaID int    `json:"marina_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newYacht := models.NewYacht(input.Name, input.MarinaID)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", r.URL.Path+"/"+strconv.Itoa(newYacht.ID))
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newYacht)
}

func UpdateYacht(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	yacht, ok := models.Yachts[id]

	if !ok {
		http.Error(w, "Yacht not found", http.StatusNotFound)
		return
	}

	ifMatch := r.Header.Get("If-Match")
	if ifMatch == "" {
		http.Error(w, "Missing If-Match header", http.StatusBadRequest)
		return
	}

	currentEtag := models.HashYacht(id)

	if ifMatch != currentEtag {
		http.Error(w, "ETag mismatch", http.StatusPreconditionFailed)
		return
	}

	var input struct {
		Name     string `json:"name"`
		MarinaID int    `json:"marina_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	yacht.Name = input.Name
	yacht.MarinaID = input.MarinaID
	models.Yachts[id] = yacht

	newEtag := models.HashYacht(id)
	w.Header().Set("Etag", newEtag)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(yacht)
}

func DeleteYacht(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	_, ok := models.Yachts[id]

	if !ok {
		http.Error(w, "Yacht not found", http.StatusNotFound)
		return
	}

	delete(models.Yachts, id)

	w.WriteHeader(http.StatusNoContent)
}
