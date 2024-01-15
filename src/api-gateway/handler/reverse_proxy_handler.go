package handler

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"sort"
)

type Service struct {
	Name         string
	ContextPath  string
	TargetURL    string
	Pattern      *regexp.Regexp
	ReverseProxy *httputil.ReverseProxy
}

type ReverseProxyHandler struct {
	Services []Service
}

// ReverseProxyConfig represents the configuration for the Reverse Proxy.
type ReverseProxyConfig struct {
	Services []Service
}

// ReverseProxyHandler is a struct representing the Reverse Proxy.

func newSingleHostReverseProxy(targetURL string) *httputil.ReverseProxy {
	target, _ := url.Parse(targetURL)
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ErrorHandler = func(writer http.ResponseWriter, request *http.Request, e error) {
		http.Error(writer, "Proxy error", http.StatusInternalServerError)
		log.Printf("proxy for %s error: %v", targetURL, e)
	}
	return proxy
}

func NewReverseProxy(config ReverseProxyConfig) *ReverseProxyHandler {
	var services []Service

	sort.Slice(config.Services, func(i, j int) bool {
		return len(config.Services[i].ContextPath) > len(config.Services[j].ContextPath)
	})

	for _, service := range config.Services {
		proxy := newSingleHostReverseProxy(service.TargetURL)
		pattern, err := regexp.Compile("^" + service.ContextPath)

		if err != nil {
			panic(err)
		}

		service.Pattern = pattern
		service.ReverseProxy = proxy

		services = append(services, service)
	}

	return &ReverseProxyHandler{
		Services: services,
	}
}

func (rp *ReverseProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming request: %s", r.URL.Path)

	for _, service := range rp.Services {
		if service.Pattern.MatchString(r.URL.Path) {
			log.Printf("Proxying %s to %s", r.URL.Path, service.Pattern.String())
			service.ReverseProxy.ServeHTTP(w, r)
			return
		}
	}

	http.Error(w, "Service not found", http.StatusNotFound)
}
