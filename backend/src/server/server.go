package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"

	"github.com/google/uuid"
	"github.com/niconielsen24/goapp/gamelogic"
)

var (
	ErrRoutesHandlersMismatch = errors.New("Routes and paths must be equally numbered")
)

type TicTacToeServer struct {
	http.Server
	games  []gamelogic.Tictactoe
	wg     sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
	mu     sync.Mutex
}

func (ts *TicTacToeServer) InitServer(addr string, rs []url.URL, hs []http.Handler) error {
	err, mux := Routes(rs, hs)
	if err != nil {
		return err
	}
	ts.Addr = addr
	ts.Handler = mux
	ctx, cancel := context.WithCancel(context.Background())
	ts.ctx = ctx
	ts.cancel = cancel

	log.Printf("Starting server at : %s\n", addr)

	if err := ts.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed while starting server : %s", err)
	}

	return nil
}

func (ts *TicTacToeServer) PopGame(id uuid.UUID) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	for i := range ts.games {
		if ts.games[i].ID == id {
			if i == len(ts.games) {
				ts.games = ts.games[:i]
			}
			ts.games = append(ts.games[:i], ts.games[i+1:]...)
			return
		}
	}
  fmt.Printf("Popped game : %v succesfully\n",id)
}

// Routes and handlers must have paired orders, otherwise a path may result in an
// undesired outcome
func Routes(rs []url.URL, hs []http.Handler) (error, *http.ServeMux) {
	mux := http.NewServeMux()

	if len(rs) != len(hs) {
		return ErrRoutesHandlersMismatch, nil
	}

	for i, route := range rs {
		mux.Handle(route.Path, CorsMiddleWare(hs[i]))
	}

	return nil, mux
}
