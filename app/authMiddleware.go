package app

import (
	"net/http"
	"strings"

	"github.com/bhaktiutama/banking/domain"
	"github.com/bhaktiutama/banking/errs"
	"github.com/gorilla/mux"
)

type AuthMiddleware struct {
	repo domain.AuthRepository
}

func (a AuthMiddleware) authorizationHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			currentRoute := mux.CurrentRoute(r)
			currentRouteVars := mux.Vars(r)
			authHeader := r.Header.Get("Authorization")

			if authHeader != "" {
				token := getTokenFromHeader(authHeader)

				isAuthorize := a.repo.IsAuthorized(token, currentRoute.GetName(), currentRouteVars)

				if isAuthorize {
					next.ServeHTTP(w, r)
				} else {
					// appError := errs.AppError{http.StatusForbidden, "Unauthorized"}

					appError := errs.AppError{Code: http.StatusForbidden, Message: "Unauthorized"}
					writeResponse(w, appError.Code, appError.AsMessage())
				}
			} else {
				writeResponse(w, http.StatusUnauthorized, "missing token")
			}
		})
	}
}

func getTokenFromHeader(header string) string {
	splitToken := strings.Split(header, "Bearer")
	if len(splitToken) == 2 {
		return strings.TrimSpace(splitToken[1])
	}
	return ""
}
