package auth

import (
	"net/http"

	"github.com/AuGhoti/goserver/helper"
	"github.com/go-chi/chi"
)

// Routes returns the chi routes for the auth endpoints
func Routes() http.Handler {
	r := chi.NewRouter()
	r.Post("/signup", signup)
	r.Post("/login", login)

	return r
}

func signup(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "hit signup endpoint"})
}

func login(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "hit login endpoint"})
}
