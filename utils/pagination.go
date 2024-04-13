package utils

import (
	"net/http"
	"strconv"

	"rest/models"
)

func Paginate(r *http.Request, data []models.Yacht) ([]models.Yacht, error) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page := 1
	limit := 10
	var err error

	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			return []models.Yacht{}, err
		}
	}

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit < 1 {
			return []models.Yacht{}, err
		}
	}

	start := (page - 1) * limit
	end := start + limit

	if start >= len(data) {
		return []models.Yacht{}, nil
	}

	if end > len(data) {
		end = len(data)
	}

	return data[start:end], nil
}
