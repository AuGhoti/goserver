package activity

import (
	"net/http"

	"github.com/AuGhoti/goserver/helper"
	"github.com/go-chi/chi"
)

// Routes sets endpoints for activities
func Routes(router chi.Router) {
	router.Route("/", func(r chi.Router) {
		r.Get("/", getAllActivities)
		r.Post("/", newActivity)
	})
	router.Route("/{id}", func(r chi.Router) {
		r.Get("/", getOneActivity)
		r.Delete("/", removeActivity)
		r.Put("/", modifyActivity)
	})
}

func getAllActivities(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "activity endpoint"})
}

func newActivity(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "new activity endpoint"})
}
func getOneActivity(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "single activity endpoint"})
}
func removeActivity(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "remove activity endpoint"})
}
func modifyActivity(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "modify activity endpoint"})
}
