package balancer

import (
	"net/http/httputil"
	"sync"
)

type Server struct {
	url          string
	ip           string
	isHealthy    bool
	mux          sync.RWMutex
	reverseProxy *httputil.ReverseProxy
}

// ApplicationLB represents an L7 balancer that acts on the application layer protocols such as HTTP, FTP, etc
type ApplicationLB struct {
	serverPool []Server
	url        string
	current    *Server // Points to the last used/current server
}

// NetworkLB represents an L4 balancer that acts on the transport layer protocols such as TCP and UDP connections
type NetworkLB struct {
	serverPool []Server
	ip         string
	current    *Server
}
