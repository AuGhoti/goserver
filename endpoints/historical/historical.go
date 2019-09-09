package historical

import (
	"net/http"

	"github.com/AuGhoti/goserver/helper"
	"github.com/go-chi/chi"
)

// Routes sets routes
func Routes(router chi.Router) {
	router.Get("/", getPastActivities)
	router.Post("/", addPastActivity)
	router.Route("/{id}", func(r chi.Router) {
		r.Get("/", getPastActivity)
		r.Delete("/", removePastActivity)
		r.Put("/", updatePastActivity)
	})
}

func getPastActivities(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "get past activities"})
}

func addPastActivity(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "add past activity"})
}

func getPastActivity(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "get past activity"})
}

func removePastActivity(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "remove past activity"})
}

func updatePastActivity(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "update past activity"})
}
