package server

import (
	"sync"
)

type Server struct {
	url       string
	ip        string
	isHealthy bool
	mux       sync.RWMutex
}

// SetAliveState will set the isHealthy property
func (s *Server) SetAliveState(alive bool) {
	s.mux.Lock()
	s.isHealthy = alive
	s.mux.Unlock()
}

// CheckLiveness will return the isHealthy property of the server
func (s *Server) CheckLiveness() bool {
	s.mux.RLock()
	state := s.isHealthy
	s.mux.RUnlock()
	return state
}
