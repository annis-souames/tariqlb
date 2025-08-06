package balancer

import (
	"fmt"
	"net/http/httputil"
	"net/url"

	"github.com/annis-souames/tariqlb/pkg/server"
)

type ApplicationLoadBalancer struct {
	url     url.URL
	proxy   *httputil.ReverseProxy
	current *server.Server
	pool    []*server.Server
}

func NewAppLoadBalancer(protocol string, host string, port int) (*ApplicationLoadBalancer, error) {
	appUrl, errUrl := url.Parse(fmt.Sprintf("%s://%s:%d", protocol, host, port))
	if errUrl != nil {
		return nil, fmt.Errorf("could not parse url string from options: %w", errUrl)
	}
	proxy := httputil.NewSingleHostReverseProxy(appUrl)
	appLb := ApplicationLoadBalancer{
		url:   *appUrl,
		proxy: proxy,
	}
	return &appLb, nil
}
func (a *ApplicationLoadBalancer) StartLB() error { return nil }
