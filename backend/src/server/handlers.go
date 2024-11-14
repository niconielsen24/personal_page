package server

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/niconielsen24/goapp/gamelogic"
)

var (
	ErrGameDoesNotExist = errors.New("Game does not exist")
)

type MoveRequest struct {
	Id       uuid.UUID          `json:"game_id"`
	Position gamelogic.Position `json:"position"`
}

type KillGameRequest struct {
	Id uuid.UUID `json:"game_id"`
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from root TicTacToeServer")
}

func InitGameHandler(ts *TicTacToeServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		t := gamelogic.Tictactoe{}
		err := t.InitTictactoe()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		ts.mu.Lock()
		defer ts.mu.Unlock()

		ts.games = append(ts.games, t)
		ts.wg.Add(1)
		setGameTimer(ts, &t)

		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(t); err != nil {
			http.Error(w, "Unable to encode game", http.StatusInternalServerError)
		}
		printRequest(r)
	}
}

func MakeMoveHandler(ts *TicTacToeServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != http.MethodPut {
			http.Error(w, "Method not allowed for this path", http.StatusMethodNotAllowed)
			return
		}
		
    defer r.Body.Close()

    var mr MoveRequest
		decode_err := json.NewDecoder(r.Body).Decode(&mr)
		if decode_err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		if validation_err := ValidateMoveRequest(ts, mr); validation_err != nil {
			http.Error(w, validation_err.Error(), http.StatusBadRequest)
		}

		game_id := mr.Id
		position := mr.Position
		game_ref := getGame(game_id, ts.games)
		if game_ref == nil {
			http.Error(w, "Game does not exist", http.StatusBadRequest)
			return
		}

		ts.mu.Lock()
		defer ts.mu.Unlock()

		if make_move_err := game_ref.MakeMove(gamelogic.PlayerX, position); make_move_err != nil {
			http.Error(w, make_move_err.Error(), http.StatusBadRequest)
			return
		}

		gamelogic.Move(game_ref)
		game_ref.GameOver()

		if encode_err := json.NewEncoder(w).Encode(*game_ref); encode_err != nil {
			http.Error(w, "Failed to encode new game state", http.StatusInternalServerError)
			return
		}
		printRequest(r)
	}
}

func KillGameHandler(ts *TicTacToeServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		defer r.Body.Close()

		var kr KillGameRequest
		if err := json.NewDecoder(r.Body).Decode(&kr); err != nil {
			http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
			return
		}
		ts.mu.Lock()
		err := killGame(&ts.games, kr.Id)
		ts.mu.Unlock()
		if err != nil {
			http.Error(w, "Failed to kill game: "+err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)

    printRequest(r)
	}
}

func printRequest(r *http.Request) {
	fmt.Printf("Request:\n")
	fmt.Printf("  Method: %s\n", r.Method)
	fmt.Printf("  URL: %s\n", r.URL.String())
	fmt.Printf("  Protocol: %s\n\n", r.Proto)

	fmt.Println("Headers:")
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Printf("  %s: %s\n", name, value)
		}
	}
	fmt.Println()

	if r.Body != nil {
		var buf bytes.Buffer
		tee := io.TeeReader(r.Body, &buf)
		bodyBytes, _ := io.ReadAll(tee)
		fmt.Printf("Body:\n%s\n", string(bodyBytes))
		r.Body = io.NopCloser(&buf)
	} else {
		fmt.Println("Body: <empty>")
	}
}

func getGame(id uuid.UUID, games []gamelogic.Tictactoe) *gamelogic.Tictactoe {
	var g *gamelogic.Tictactoe
	for i := range games {
		if games[i].ID == id {
			g = &games[i]
			break
		}
	}
	return g
}

func setGameTimer(ts *TicTacToeServer, t *gamelogic.Tictactoe) {
	go func(ctx context.Context, t *gamelogic.Tictactoe) {
		defer ts.wg.Done()
		select {
		case <-time.After(20 * time.Minute):
			ts.PopGame(t.ID)
			fmt.Println("Game removed: ", t.ID)
		case <-ctx.Done():
			fmt.Println("Goroutine cancelled for game:", t.ID)
		}
	}(ts.ctx, t)
}

func killGame(games *[]gamelogic.Tictactoe, g uuid.UUID) error {
	for i := range *games {
		if (*games)[i].ID == g {
			*games = append((*games)[:i], (*games)[i+1:]...)
			return nil
		}
	}
	return ErrGameDoesNotExist
}
