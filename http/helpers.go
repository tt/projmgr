package main

import (
	"encoding/json"
	"net/http"

	. "github.com/heroku/projmgr"
)

func readJSON(r *http.Request, dest interface{}) error {
	return json.NewDecoder(r.Body).Decode(dest)
}

func writeJSON(w http.ResponseWriter, body interface{}) {
	json.NewEncoder(w).Encode(body)
}

func writeError(w http.ResponseWriter, err error) {
	switch err {
	case ErrNotFound:
		http.Error(w, "not found", http.StatusNotFound)
	default:
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
