package handlers_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"rest/helpers"
	"rest/models"
)

func TestListMarinas(t *testing.T) {
	models.ClearMarinas()
	models.NewMarina("Test Marina 1")
	models.NewMarina("Test Marina 2")

	rr := helpers.GetRequest(t, "/marinas")

	var marinas []models.Marina
	assert.Nil(t, json.NewDecoder(rr.Body).Decode(&marinas))
	assert.Equal(t, 2, len(marinas))
}

func TestGetMarina(t *testing.T) {
	testMarina := models.NewMarina("Test Marina")

	rr := helpers.GetRequest(t, fmt.Sprintf("/marinas/%d", testMarina.ID))

	var marina models.Marina
	assert.Nil(t, json.NewDecoder(rr.Body).Decode(&marina))
	assert.Equal(t, testMarina.Name, marina.Name)
}

func TestCreateMarina(t *testing.T) {
	input := struct {
		Name string `json:"name"`
	}{
		Name: "Test Marina",
	}

	body, _ := json.Marshal(input)

	rr := helpers.CreateRequest(t, "/marinas", string(body))
	assert.Equal(t, 201, rr.Code)

	var marina models.Marina
	assert.Nil(t, json.NewDecoder(rr.Body).Decode(&marina))
	assert.Equal(t, input.Name, marina.Name)
	assert.Equal(t, models.Marinas[marina.ID].ID, marina.ID)
}

func TestUpdateMarina(t *testing.T) {
	testMarina := models.NewMarina("Test Marina")

	input := struct {
		Name string `json:"name"`
	}{
		Name: "Updated Test Marina",
	}

	body, _ := json.Marshal(input)

	rr := helpers.UpdateRequest(t, fmt.Sprintf("/marinas/%d", testMarina.ID), string(body))

	var marina models.Marina
	assert.Nil(t, json.NewDecoder(rr.Body).Decode(&marina))
	assert.Equal(t, input.Name, marina.Name)
	assert.Equal(t, models.Marinas[marina.ID].Name, input.Name)
}

func TestDeleteMarina(t *testing.T) {
	testMarina := models.NewMarina("Test Marina")

	rr := helpers.DeleteRequest(t, fmt.Sprintf("/marinas/%d", testMarina.ID))

	assert.Equal(t, "", rr.Body.String())
}

func TestListYachtsInMarina(t *testing.T) {
	marina := models.NewMarina("Test Marina")
	models.NewYacht("Test Yacht 1", marina.ID)
	models.NewYacht("Test Yacht 2", marina.ID)

	unusedMarina := models.NewMarina("Test Marina")
	models.NewYacht("Test Yacht 3", unusedMarina.ID)
	models.NewYacht("Test Yacht 4", unusedMarina.ID)

	rr := helpers.GetRequest(t, fmt.Sprintf("/marinas/%d/yachts", marina.ID))

	var yachts []models.Yacht
	assert.Nil(t, json.NewDecoder(rr.Body).Decode(&yachts))
	assert.Equal(t, 2, len(yachts))
}
