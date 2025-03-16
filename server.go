package obstrat

import (
	"fmt"
	"net/http"
)

type StratStore interface {
	GetRandomStrat() string
	ShuffleStratPile()
	DrawTopStrat() string
}

type StratServer struct {
	store StratStore
	http.Handler
}

func NewStratServer(store StratStore) *StratServer {
	s := new(StratServer)
	s.store = store

	router := http.NewServeMux()

	router.Handle("/next-strategy", http.HandlerFunc(s.stratHandler))
	router.Handle("/shuffle", http.HandlerFunc(s.shuffleHandler))
	s.Handler = router
	return s
}

func (s *StratServer) stratHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		s.fetchStrat(w)
	}

}

func (s *StratServer) fetchStrat(w http.ResponseWriter) {
	strat := s.store.DrawTopStrat()

	fmt.Fprint(w, strat)
}

func (s *StratServer) shuffleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		s.store.ShuffleStratPile()
	}
	fmt.Fprint(w, "deck shuffled, obliquely")
}
