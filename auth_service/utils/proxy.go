package utils

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func ProxyToService(targetBaseURL string, pathPrefix string) http.HandlerFunc {

	target, err := url.Parse(targetBaseURL)
	if err != nil {
		fmt.Println("Error parsing target URL", err)
		return nil
	}

	// Where to forward the request, basically backend service ka hostname 
	proxy := httputil.NewSingleHostReverseProxy(target)

	originalDirector := proxy.Director

	// Override the default Director function of the reverse proxy.
// The Director modifies the incoming request before it's sent to the backend.
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		originalPath := req.URL.Path

		strippedPath := strings.TrimPrefix(originalPath, pathPrefix)

		// Sets host and path that the request has to be forwarded to
		req.URL.Host = target.Host
		req.URL.Path = target.Path + strippedPath
		req.Host = target.Host

		if userId, ok := req.Context().Value("userID").(string); ok {
			req.Header.Set("X-User-ID", userId)
		}
	}

	return proxy.ServeHTTP
}