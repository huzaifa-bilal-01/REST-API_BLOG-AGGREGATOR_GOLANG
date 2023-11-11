package main

import (
	"fmt"
	"net/http"

	"github.com/huzaifa-bilal-01/rss-server/internal/auth"
	"github.com/huzaifa-bilal-01/rss-server/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIkey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error %v", err))
			return
		}
		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't get User %v", err))
			return
		}

		handler(w, r, user)
	}
}
