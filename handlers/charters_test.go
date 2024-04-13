package handlers_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"rest/helpers"
	"rest/models"
)

func TestListCharters(t *testing.T) {
	marina := models.NewMarina("Test Marina")
	models.NewCharter("Test Charter 1", marina.ID)
	models.NewCharter("Test Charter 2", marina.ID)

	rr := helpers.GetRequest(t, "/charters")

	var charters []models.Charter
	assert.Nil(t, json.NewDecoder(rr.Body).Decode(&charters))
	assert.Equal(t, 2, len(charters))
}

func TestGetCharter(t *testing.T) {
	marina := models.NewMarina("Test Marina")
	testCharter := models.NewCharter("Test Charter", marina.ID)

	rr := helpers.GetRequest(t, fmt.Sprintf("/charters/%d", testCharter.ID))

	var charter models.Charter
	assert.Nil(t, json.NewDecoder(rr.Body).Decode(&charter))
	assert.Equal(t, testCharter.Captain, charter.Captain)
}

func TestCreateCharter(t *testing.T) {
	marina := models.NewMarina("Test Marina")

	input := struct {
		Captain string `json:"captain"`
		YachtID int    `json:"yacht_id"`
	}{
		Captain: "Test Charter",
		YachtID: marina.ID,
	}

	body, _ := json.Marshal(input)

	rr := helpers.CreateRequest(t, "/charters", string(body))
	assert.Equal(t, 201, rr.Code)

	var charter models.Charter
	assert.Nil(t, json.NewDecoder(rr.Body).Decode(&charter))
	assert.Equal(t, input.Captain, charter.Captain)
	assert.Equal(t, input.YachtID, charter.YachtID)
	assert.Equal(t, models.Charters[charter.ID].ID, charter.ID)
}

func TestUpdateCharter(t *testing.T) {
	marina := models.NewMarina("Test Marina")
	testCharter := models.NewCharter("Test Charter", marina.ID)

	secondMarina := models.NewMarina("Test Marina 2")

	input := struct {
		Captain string `json:"captain"`
		YachtID int    `json:"yacht_id"`
	}{
		Captain: "Updated Test Charter",
		YachtID: secondMarina.ID,
	}

	body, _ := json.Marshal(input)

	rr := helpers.UpdateRequest(t, fmt.Sprintf("/charters/%d", testCharter.ID), string(body))

	var charter models.Charter
	assert.Nil(t, json.NewDecoder(rr.Body).Decode(&charter))
	assert.Equal(t, input.Captain, charter.Captain)
	assert.Equal(t, input.YachtID, charter.YachtID)
	assert.Equal(t, models.Charters[charter.ID].Captain, input.Captain)
}

func TestDeleteCharter(t *testing.T) {
	marina := models.NewMarina("Test Marina")
	testCharter := models.NewCharter("Test Charter", marina.ID)

	rr := helpers.DeleteRequest(t, fmt.Sprintf("/charters/%d", testCharter.ID))

	assert.Equal(t, "", rr.Body.String())
}
