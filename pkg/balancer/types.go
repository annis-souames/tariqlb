package balancer

import (
	"net/http/httputil"
	"sync"
)

type Server struct {
	url       string
	ip        string
	isHealthy bool
	mux       sync.RWMutex
}

// ApplicationLB is a load balancer that acts on L7 protocols: HTTP and HTTPs
type ApplicationLoadBalancer struct {
	serverPool []Server
	url        string
	current    *Server // Points to the last used/current server
	proxy      *httputil.ReverseProxy
}
