package server

import (
	"errors"
	"log"
	"net/http"
	"net/url"

  "github.com/niconielsen24/goapp/gamelogic"
)

var (
	ErrRoutesHandlersMismatch = errors.New("Routes and paths must be equally numbered")
)

type TicTacToeServer struct {
	http.Server
  games []gamelogic.Tictactoe
}

func (ts *TicTacToeServer) InitServer(addr string ,rs []url.URL, hs []http.Handler) error {
	err, mux := Routes(rs, hs)
	if err != nil {
		return err
	}
  ts.Addr = addr
  ts.Handler = mux

  log.Printf("Starting server at : %s\n", addr)

  if err := ts.ListenAndServe(); err != nil && err != http.ErrServerClosed {
    log.Fatalf("Failed while starting server : %s", err)
  } 

  return nil
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}


// Routes and handlers must have paired orders, otherwise a path may result in an
// undesired outcome
func Routes(rs []url.URL, hs []http.Handler) (error, *http.ServeMux) {
	mux := http.NewServeMux()

	if len(rs) != len(hs) {
		return ErrRoutesHandlersMismatch, nil
	}

	for i, route := range rs {
		mux.Handle(route.Path, hs[i])
	}

	return nil, mux
}
