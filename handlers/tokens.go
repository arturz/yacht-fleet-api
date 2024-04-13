package handlers

import (
	"encoding/json"
	"net/http"
	"rest/models"
	"strconv"
)

func CreateToken(w http.ResponseWriter, r *http.Request) {
	token := models.NewToken()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", r.URL.Path+"/"+strconv.Itoa(token.ID))
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(token)
}
