package auth

import (
	"net/http"

	"github.com/urfave/negroni"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var userSession *session
		var err error

		if userSession, err = getSessionToken(r); err != nil {
			http.Error(w, "Please login", http.StatusUnauthorized)
			return
		}

		lrw := negroni.NewResponseWriter(w)
		next.ServeHTTP(lrw, r)

		if r.URL.Path == loginMethod && lrw.Status() == http.StatusOK {
			createSession(lrw, userSession)
		}
	})
}
