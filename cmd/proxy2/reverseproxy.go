package proxy2

import (
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewMultipleHostsReverseProxy(targets []*url.URL) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		target := targets[rand.Int()%len(targets)]
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = target.Path
	}
	return &httputil.ReverseProxy{Director: director}
}
func Main() {
	proxy := NewMultipleHostsReverseProxy([]*url.URL{
		{
			Scheme: "https",
			Host:   "www.sdinc.cn",
		},
		{
			Scheme: "https",
			Host:   "ci.sdinc.cn",
		},
	})
	log.Fatal(http.ListenAndServe(":9090", proxy))
}
