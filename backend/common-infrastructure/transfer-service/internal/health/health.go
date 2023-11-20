package health

import (
	"fmt"
	"net/http"

	routing "github.com/gorilla/mux"
)

func RegisterHandlers(r *routing.Router) {
	r.HandleFunc("/health", healthHandler).Methods("GET")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}
