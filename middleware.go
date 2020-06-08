package jwt

import (
	"context"
	"net/http"
)

// Authenticator verifies authentication provided in the request's header
func Authenticator(t *Token) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, req *http.Request) {
			tokenString, err := t.GetToken(req)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			data, err := t.CheckToken(tokenString)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			// If we get here, everything worked and we can set the
			// user data in context
			newReq := req.WithContext(context.WithValue(req.Context(), t.options.UserProperty, data))
			// update the current request with the new context information
			*req = *newReq
			next.ServeHTTP(w, req)
		}
		return http.HandlerFunc(fn)
	}
}
