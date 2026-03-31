package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func MainEndpointHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "MainEndpoint: %v\n", vars["category"])
}
