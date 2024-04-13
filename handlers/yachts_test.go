package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"rest/helpers"
	"rest/models"
)

func TestListYachts(t *testing.T) {
	models.ClearYachts()
	marina := models.NewMarina("Test Marina")
	models.NewYacht("Test Yacht 1", marina.ID)
	models.NewYacht("Test Yacht 2", marina.ID)

	rr := helpers.GetRequest(t, "/yachts")

	var yachts []models.Yacht
	assert.Nil(t, json.NewDecoder(rr.Body).Decode(&yachts))
	assert.Equal(t, 2, len(yachts))
}

func TestListYachtsPagination(t *testing.T) {
	models.ClearYachts()
	marina := models.NewMarina("Test Marina")
	models.NewYacht("Test Yacht 1", marina.ID)
	models.NewYacht("Test Yacht 2", marina.ID)
	models.NewYacht("Test Yacht 3", marina.ID)
	models.NewYacht("Test Yacht 4", marina.ID)
	models.NewYacht("Test Yacht 5", marina.ID)

	rr := helpers.GetRequest(t, "/yachts?limit=2&page=2")

	var yachts []models.Yacht
	assert.Nil(t, json.NewDecoder(rr.Body).Decode(&yachts))
	assert.Equal(t, 2, len(yachts))
	assert.Equal(t, "Test Yacht 3", yachts[0].Name)
	assert.Equal(t, "Test Yacht 4", yachts[1].Name)
}

func TestGetYacht(t *testing.T) {
	marina := models.NewMarina("Test Marina")
	testYacht := models.NewYacht("Test Yacht", marina.ID)

	rr := helpers.GetRequest(t, fmt.Sprintf("/yachts/%d", testYacht.ID))

	var yacht models.Yacht
	assert.Nil(t, json.NewDecoder(rr.Body).Decode(&yacht))
	assert.Equal(t, testYacht.Name, yacht.Name)

	etag := models.HashYacht(testYacht.ID)
	assert.Equal(t, etag, rr.Header().Get("ETag"))
}

func TestCreateYacht(t *testing.T) {
	marina := models.NewMarina("Test Marina")

	input := struct {
		Name     string `json:"name"`
		MarinaID int    `json:"marina_id"`
	}{
		Name:     "Test Yacht",
		MarinaID: marina.ID,
	}

	body, _ := json.Marshal(input)

	rr := helpers.CreateRequest(t, "/yachts", string(body))
	assert.Equal(t, 201, rr.Code)

	var yacht models.Yacht
	assert.Nil(t, json.NewDecoder(rr.Body).Decode(&yacht))
	assert.Equal(t, input.Name, yacht.Name)
	assert.Equal(t, input.MarinaID, yacht.MarinaID)
	assert.Equal(t, models.Yachts[yacht.ID].ID, yacht.ID)
}
func TestUpdateYacht(t *testing.T) {
	marina := models.NewMarina("Test Marina")
	testYacht := models.NewYacht("Test Yacht", marina.ID)

	secondMarina := models.NewMarina("Test Marina 2")

	input := struct {
		Name     string `json:"name"`
		MarinaID int    `json:"marina_id"`
	}{
		Name:     "Updated Test Yacht",
		MarinaID: secondMarina.ID,
	}

	body, _ := json.Marshal(input)

	initialEtag := models.HashYacht(testYacht.ID)

	headers := map[string]string{
		"If-Match":     initialEtag,
		"Content-Type": "application/json",
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/yachts/%d", testYacht.ID), strings.NewReader(string(body)))
	assert.Nil(t, err)

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// pierwsze zapytanie
	rr := helpers.UpdateRequestWithHeaders(t, req)

	var yacht models.Yacht
	assert.Nil(t, json.NewDecoder(rr.Body).Decode(&yacht))
	assert.Equal(t, input.Name, yacht.Name)
	assert.Equal(t, input.MarinaID, yacht.MarinaID)
	assert.Equal(t, rr.Code, http.StatusOK)

	req, err = http.NewRequest(http.MethodPut, fmt.Sprintf("/yachts/%d", testYacht.ID), strings.NewReader(string(body)))
	assert.Nil(t, err)

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// drugie zapytanie z tym samym ETagiem
	rr = helpers.UpdateRequestWithHeaders(t, req)

	assert.Equal(t, http.StatusPreconditionFailed, rr.Code)
}

func TestDeleteYacht(t *testing.T) {
	marina := models.NewMarina("Test Marina")
	testYacht := models.NewYacht("Test Yacht", marina.ID)

	rr := helpers.DeleteRequest(t, fmt.Sprintf("/yachts/%d", testYacht.ID))

	assert.Equal(t, "", rr.Body.String())
}
