package helper

import (
	"encoding/json"
	"net/http"
)

// RespondWithJSON handles responding with JSON to reduce redundant code
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
