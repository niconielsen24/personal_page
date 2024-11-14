package server

import (
	"testing"

  "github.com/stretchr/testify/assert"
	"github.com/google/uuid"
	"github.com/niconielsen24/goapp/gamelogic"
)

func testGetGame(t *testing.T) {
	var mockId uuid.UUID = uuid.New()
	var mockGames = []gamelogic.Tictactoe{
		{
			ID: mockId,
		},
	}
	if pointer := getGame(mockId, mockGames); pointer == nil {
		t.Errorf("expected game id to exist got: %v", pointer)
	}

	mockId = uuid.Nil

	if pointer := getGame(mockId, mockGames); pointer != nil {
		t.Errorf("expected game id to not exist got: %v", pointer)
	}
}

func TestKillGame(t *testing.T) {
	mockID1 := uuid.New()
	mockID2 := uuid.New()

	mockGames := []gamelogic.Tictactoe{
		{ID: mockID1},
		{ID: mockID2},
	}

	// Test case: Game exists and should be removed
	err := killGame(&mockGames, mockID1)
	assert.Nil(t, err)
	assert.Len(t, mockGames, 1)
	assert.Equal(t, mockGames[0].ID, mockID2)

	// Test case: Game does not exist
	err = killGame(&mockGames, uuid.New())
	assert.ErrorIs(t, err, ErrGameDoesNotExist)
}
