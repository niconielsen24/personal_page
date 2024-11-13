package server

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/niconielsen24/goapp/gamelogic"
)

var (
	ErrInvalidUUID      = errors.New("Invalid uuid")
	ErrUuidDoesNotExist = errors.New("Uuid does not exist")
	ErrInvalidPosition  = errors.New("Position is invalid")
	ErrGameIsOver       = errors.New("Game has ended")
)

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func CorsMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		EnableCors(&w)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func ValidateMoveRequest(ts *TicTacToeServer, mr MoveRequest) error {
	if mr.Id == uuid.Nil {
		return ErrInvalidUUID
	}
	res := false
	for i := range ts.games {
		res = res || mr.Id == ts.games[i].ID
	}
	if !res {
		return ErrUuidDoesNotExist
	}
	if res {
		for i := range ts.games {
			if ts.games[i].ID == mr.Id {
				if ts.games[i].Winner != gamelogic.Empty {
					return ErrGameIsOver
				}
			}
		}
	}
	if mr.Position.OutOfBounds() {
		return ErrInvalidPosition
	}

	return nil
}
