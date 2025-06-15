package cors

import (
	"net/http"
	"slices"
	"strings"
)

var originAllowlist = []string{
	"http://localhost:5173",
	"https://runquanrayzhou.netlify.app/",
	"https://linknyc-finder.netlify.app/",
	"https://quiz-me-trivia-app.netlify.app/",
	"https://pocket-dictionary-app.netlify.app/",
}

var methodAllowlist = []string{"GET", "POST", "DELETE", "OPTIONS"}

func isPreflight(r *http.Request) bool {
	return r.Method == "OPTIONS" &&
		r.Header.Get("Origin") != "" &&
		r.Header.Get("Access-Control-Request-Method") != ""
}

func CheckCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isPreflight(r) {
			origin := r.Header.Get("Origin")
			method := r.Header.Get("Access-Control-Request-Method")
			// currently only allowed for own websites and local testing
			if slices.Contains(originAllowlist, origin) && slices.Contains(methodAllowlist, method) {
				// w.Header().Set("Access-Control-Allow-Origin", "*") For all origins
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", strings.Join(methodAllowlist, ", "))
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}
			w.Header().Add("Vary", "Origin")
			next.ServeHTTP(w, r)
		}
		// Not a preflight: regular request.
		origin := r.Header.Get("Origin")
		if slices.Contains(originAllowlist, origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}
		w.Header().Add("Vary", "Origin")
		next.ServeHTTP(w, r)
	})
}
