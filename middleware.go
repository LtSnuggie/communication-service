package communication

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"method":  r.Method,
			"path":    r.RequestURI,
			"date":    time.Now(),
			"token":   r.Header.Get("Authorization"),
			"srcaddr": r.RemoteAddr,
		}).Info("request received")
		// r.Body.Close()
		next.ServeHTTP(w, r)
	})
}

func GetAuthMiddleware(noauth map[string][]string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, method := range noauth[r.URL.Path] {
				if method == r.Method {
					next.ServeHTTP(w, r)
					return
				}
			}
			//DO AUTH
			next.ServeHTTP(w, r)
		})
	}
}
