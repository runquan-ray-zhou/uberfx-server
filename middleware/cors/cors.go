package cors

import (
	"net/http"
	"slices"
)

var originAllowlist = []string{
	"http://localhost:5173",
	"https://runquanrayzhou.netlify.app/",
	"https://linknyc-finder.netlify.app/",
	"https://quiz-me-trivia-app.netlify.app/",
	"https://pocket-dictionary-app.netlify.app/",
}

func CheckCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if slices.Contains(originAllowlist, origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		w.Header().Add("Vary", "Origin")
		next.ServeHTTP(w, r)
	})
}
