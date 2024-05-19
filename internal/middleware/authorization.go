package middleware

import (
	"errors"
	"net/http"

	"github.com/aneesazc/goapi/api"
	"github.com/aneesazc/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Unauthorized")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization Header
		var username string = r.URL.Query().Get("username")
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || username == "" {
			log.Error("Authorization Header or username can't be missing")
			tools.WriteError(w, api.Error{Code: http.StatusUnauthorized, Message: UnAuthorizedError.Error()})
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		if loginDetails == nil || (authHeader != (*loginDetails).AuthToken) {
			log.Error("Unauthorized")
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
