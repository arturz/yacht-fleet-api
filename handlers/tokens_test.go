package handlers_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"rest/helpers"
	"rest/models"
	"testing"
)

func TestCreateToken(t *testing.T) {
	rr := helpers.CreateRequest(t, "/tokens", "")

	var token models.Token
	if err := json.NewDecoder(rr.Body).Decode(&token); err != nil {
		t.Errorf("error decoding response body: %v", err)
	}

	assert.Equal(t, 201, rr.Code)
	assert.Equal(t, false, models.IsTokenUsed(token.ID))

	models.UseToken(token.ID)

	assert.Equal(t, true, models.IsTokenUsed(token.ID))
}
