package balancer

import "sync"

type Server struct {
	Url       string
	IsHealthy bool
	Mux       sync.RWMutex
}

type LoadBalancer struct {
	ServerPool []Server
	URL        string
	Current    *Server // Points to the last used/current server
}
