package handlers_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"rest/helpers"
	"rest/models"
)

func TestListMigrations(t *testing.T) {
	models.ClearMigrations()
	oldMarina := models.NewMarina("Test Marina")
	newMarina := models.NewMarina("Test Marina 2")
	yacht := models.NewYacht("Test Yacht", oldMarina.ID)
	models.NewMigration(yacht.ID, newMarina.ID)

	rr := helpers.GetRequest(t, "/migrations")

	var migrations []models.Migration
	assert.Nil(t, json.NewDecoder(rr.Body).Decode(&migrations))
	assert.Equal(t, 1, len(migrations))
	assert.Equal(t, yacht.ID, migrations[0].YachtID)
	assert.Equal(t, newMarina.ID, migrations[0].MarinaID)
	assert.Equal(t, oldMarina.ID, migrations[0].SourceMarinaID)
}

func TestGetMigration(t *testing.T) {
	models.ClearMigrations()
	marina := models.NewMarina("Test Marina")
	yacht := models.NewYacht("Test Yacht", marina.ID)
	migration, err := models.NewMigration(yacht.ID, marina.ID)
	assert.Nil(t, err)

	rr := helpers.GetRequest(t, fmt.Sprintf("/migrations/%d", migration.ID))

	var m models.Migration
	assert.Nil(t, json.NewDecoder(rr.Body).Decode(&m))
	assert.Equal(t, migration.ID, m.ID)
}

func TestCreateMigration(t *testing.T) {
	oldMarina := models.NewMarina("Test Marina")
	newMarina := models.NewMarina("Test Marina 2")
	yacht := models.NewYacht("Test Yacht", oldMarina.ID)

	rr := helpers.CreateRequest(t, "/tokens", "")
	assert.Equal(t, 201, rr.Code)

	var token models.Token
	assert.Nil(t, json.NewDecoder(rr.Body).Decode(&token))

	input := struct {
		YachtID  int `json:"yacht_id"`
		MarinaID int `json:"marina_id"`
	}{
		YachtID:  yacht.ID,
		MarinaID: newMarina.ID,
	}

	body, _ := json.Marshal(input)

	rr = helpers.CreateRequest(t, fmt.Sprintf("/migrations?token=%d", token.ID), string(body))

	var migration models.Migration
	assert.Nil(t, json.NewDecoder(rr.Body).Decode(&migration))
	assert.Equal(t, 201, rr.Code)
	assert.Equal(t, yacht.ID, migration.YachtID)
	assert.Equal(t, oldMarina.ID, migration.SourceMarinaID)
	assert.Equal(t, newMarina.ID, migration.MarinaID)

	// sprobuj znowu z tym samym tokenem

	rr = helpers.CreateRequest(t, fmt.Sprintf("/migrations?token=%d", token.ID), string(body))

	assert.Equal(t, 400, rr.Code)

	// sprobuj znowu z nowym samym tokenem

	rr = helpers.CreateRequest(t, "/tokens", "")
	assert.Nil(t, json.NewDecoder(rr.Body).Decode(&token))

	input = struct {
		YachtID  int `json:"yacht_id"`
		MarinaID int `json:"marina_id"`
	}{
		YachtID:  yacht.ID,
		MarinaID: oldMarina.ID,
	}

	body, _ = json.Marshal(input)

	rr = helpers.CreateRequest(t, fmt.Sprintf("/migrations?token=%d", token.ID), string(body))
	assert.Nil(t, json.NewDecoder(rr.Body).Decode(&migration))
	assert.Equal(t, 201, rr.Code)
	assert.Equal(t, yacht.ID, migration.YachtID)
	assert.Equal(t, newMarina.ID, migration.SourceMarinaID)
	assert.Equal(t, oldMarina.ID, migration.MarinaID)
}
