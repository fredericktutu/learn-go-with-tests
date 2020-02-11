package main

import (
	"fmt"
	"net/http"
	"strings"
)

// the database has a method to get score
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

//Server includes a database
type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) processWin(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
	return
}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	var score int
	if score = p.store.GetPlayerScore(player); score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, r)
	case http.MethodGet:
		p.showScore(w, r)
	}
}

func GetPlayerScore(name string) int {
	if name == "Pepper" {
		return 20
	}

	if name == "Floyd" {
		return 10
	}

	return 0

}
