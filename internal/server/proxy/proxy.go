package proxy

import (
	"go.opencensus.io/plugin/ochttp"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func ListenProxy(apiServerPort, staticServerPort, proxyServerPort string) error {
	apiServerPort = "http://localhost" + apiServerPort
	staticServerPort = "http://localhost" + staticServerPort

	apiServerUrl, err := url.Parse(apiServerPort)
	if err != nil {
		return err
	}

	staticServerUrl, err := url.Parse(staticServerPort)
	if err != nil {
		return err
	}

	apiServerReverseProxy := httputil.NewSingleHostReverseProxy(apiServerUrl)
	apiServerReverseProxy.Transport = &ochttp.Transport{}

	staticServerReverseProxy := httputil.NewSingleHostReverseProxy(staticServerUrl)
	staticServerReverseProxy.Transport = &ochttp.Transport{}

	f := http.NewServeMux()
	f.HandleFunc("/", proxyHandler(staticServerReverseProxy, apiServerReverseProxy))

	log.Println("Listening proxy server on ", proxyServerPort)
	err = http.ListenAndServe(proxyServerPort, f)
	if err != nil {
		return err
	}
	return nil
}

func proxyHandler(staticServerReverseProxy, apiServerReverseProxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Ben", "Rad")

		if strings.Contains(r.URL.Path, "/api") {
			apiServerReverseProxy.ServeHTTP(w, r)
		} else {
			staticServerReverseProxy.ServeHTTP(w, r)
		}
	}
}
