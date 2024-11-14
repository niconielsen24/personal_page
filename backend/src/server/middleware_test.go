package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/niconielsen24/goapp/gamelogic"
)

func TestCorsMiddleWare(t *testing.T) {
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	wrappedHandler := CorsMiddleWare(mockHandler)

	// Create a request to test.
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	// Call the wrapped handler.
	wrappedHandler.ServeHTTP(w, req)

	// Verify the response headers.
	resp := w.Result()
	defer resp.Body.Close()

	if origin := resp.Header.Get("Access-Control-Allow-Origin"); origin != "*" {
		t.Errorf("expected Access-Control-Allow-Origin to be '*', got '%s'", origin)
	}

	if methods := resp.Header.Get("Access-Control-Allow-Methods"); methods == "" {
		t.Errorf("expected Access-Control-Allow-Methods header, got none")
	}

	// Verify the response body for non-OPTIONS requests.
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code 200, got %d", resp.StatusCode)
	}
}

func TestValidateMoveRequest(t *testing.T) {
	var mockId uuid.UUID = uuid.New()
	mockReq := MoveRequest{
		Id: mockId,
		Position: gamelogic.Position{
			X: 0,
			Y: 0,
		}}

	var mockServer = TicTacToeServer{
		games: []gamelogic.Tictactoe{
			{
				ID: mockId,
			},
		}}
	if err := ValidateMoveRequest(&mockServer, mockReq); err != ErrGameIsOver {
		t.Errorf("expected game has ended, got %v", err)
	}

	mockServer.games[0].Winner = gamelogic.Empty

	if err := ValidateMoveRequest(&mockServer, mockReq); err != nil {
		t.Errorf("expected move request to be valid, got %v", err)
	}

	mockReq.Id = uuid.Nil

	if err := ValidateMoveRequest(&mockServer, mockReq); err != ErrInvalidUUID {
		t.Errorf("expected request id to be invalid, got %v", err)
	}

	mockReq.Id = mockId
	mockServer.games[0].ID = uuid.New()

	if err := ValidateMoveRequest(&mockServer, mockReq); err != ErrUuidDoesNotExist {
		t.Errorf("expected uuid does not exist, got %v", err)
	}
}
