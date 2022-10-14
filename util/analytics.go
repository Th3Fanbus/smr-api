package util

import (
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func HandleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	serveReverseProxy("https://www.google-analytics.com", res, req)
}

func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	redirectURL, _ := url.Parse(target)

	proxy := httputil.NewSingleHostReverseProxy(redirectURL)

	req.URL.Host = redirectURL.Host
	req.URL.Scheme = redirectURL.Scheme

	query := req.URL.Query()
	query.Set("uip", RealIP(req))

	req.URL.RawQuery = query.Encode()

	req.URL.Path = strings.ReplaceAll(req.URL.Path, "analytics/cozzect", "collect")
	req.URL.Path = strings.ReplaceAll(req.URL.Path, "analytics/r/cozzect", "r/collect")
	req.URL.Path = strings.ReplaceAll(req.URL.Path, "analytics/j/cozzect", "j/collect")

	req.RequestURI = strings.ReplaceAll(req.RequestURI, "analytics/cozzect", "collect")
	req.RequestURI = strings.ReplaceAll(req.RequestURI, "analytics/r/cozzect", "r/collect")
	req.RequestURI = strings.ReplaceAll(req.RequestURI, "analytics/j/cozzect", "j/collect")

	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = redirectURL.Host

	proxy.ServeHTTP(res, req)
}

func RealIP(req *http.Request) string {
	header := req.Header

	if ip := header.Get("X-Forwarded-For"); ip != "" {
		return strings.Split(ip, ", ")[0]
	}

	if ip := header.Get("X-Real-IP"); ip != "" {
		return ip
	}

	ra, _, _ := net.SplitHostPort(req.RemoteAddr)

	return ra
}
