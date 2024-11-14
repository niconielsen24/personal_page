package main

import (
	"net/http"
	"net/url"

	"github.com/niconielsen24/goapp/server"
)

func main() {
	ts := &server.TicTacToeServer{}
	routes := []url.URL{
		{Path: "/"},
		{Path: "/init"},
    {Path: "/makeMove"},
    {Path: "/killGame"},
	}

	handlers := []http.Handler{
		http.HandlerFunc(server.RootHandler),
		http.HandlerFunc(server.InitGameHandler(ts)),
    http.HandlerFunc(server.MakeMoveHandler(ts)),
    http.HandlerFunc(server.KillGameHandler(ts)),
	}

	ts.InitServer(":8000", routes, handlers)
}
