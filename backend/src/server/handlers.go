package server

import (
	"encoding/json"
	"fmt"
	"net/http"
  "bytes"
  "io"
  
	"github.com/google/uuid"
	"github.com/niconielsen24/goapp/gamelogic"
)

type MoveRequest struct {
	Id       uuid.UUID          `json:"game_id"`
	Position gamelogic.Position `json:"position"`
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

		ts.games = append(ts.games, t)

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
		var game_ref *gamelogic.Tictactoe
		for i := range ts.games {
			if ts.games[i].ID == game_id {
				game_ref = &ts.games[i]
				break
			}
		}
		if game_ref == nil {
			http.Error(w, "Game does not exist", http.StatusBadRequest)
			return
		}
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
