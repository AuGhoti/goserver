package current

import (
	"net/http"

	"github.com/AuGhoti/goserver/helper"
	"github.com/go-chi/chi"
)

// Routes sets current routes
func Routes(router chi.Router) {
	router.Route("/", func(r chi.Router) {
		r.Get("/", getCurrent)
		r.Post("/", addPastActivity)
	})
	// TODO: with /{id} /start gets skipped
	router.Post("/start", startActivity)
	router.Get("/end/{id}", stopActivity)
	router.Route("/{id}", func(r chi.Router) {
		r.Get("/", getOneActivity)
		r.Put("/", updateActivity)
		r.Delete("/", deleteActivity)
	})
}

func getCurrent(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "get current timers"})
}
func addPastActivity(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "add past activity"})
}
func startActivity(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "start timer"})
}
func stopActivity(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "stop timer"})
}
func getOneActivity(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "get one timers"})
}
func updateActivity(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "update timer"})
}
func deleteActivity(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "delete timer"})
}
