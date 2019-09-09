package stats

import (
	"net/http"

	"github.com/AuGhoti/goserver/helper"
	"github.com/go-chi/chi"
)

// Routes sets routes
func Routes(route chi.Router) {
	route.Get("/", getStats)
}

func getStats(w http.ResponseWriter, r *http.Request) {
	helper.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "get some stats"})
}
