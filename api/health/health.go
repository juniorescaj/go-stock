package health

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Map(r *mux.Router) {
	r.HandleFunc("/health", get)
}

func get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
