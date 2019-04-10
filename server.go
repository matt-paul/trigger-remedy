package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}
	r := mux.NewRouter()
	// If an incoming request URL matches one of the paths, the corresponding handler is
	// called, passing (http.ResponseWriter, *http.Request) as parameters
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/remedies", RemediesHandler)
	r.HandleFunc("/triggers", TriggersHandler)
	r.HandleFunc("/remedies/{category}/", RemediesCategoryHandler)
	// Paths can have variables, defined using the format {name} or {name:pattern}
	// r.HandleFunc("/triggers/{id:[0-9]+}", TriggerHandler)
	http.Handle("/", r)

	http.ListenAndServe(":3001", r)
}

// HomeHandler handlers the "/" endpoint
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	x := "world"
	fmt.Fprintf(w, "Fantastic stuff hello %v\n", x)
}

// RemediesHandler handles the "/remedies" endpoint
func RemediesHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	x := "These are you're remedies"
	fmt.Fprintf(w, "Right: %v\n", x)
}

// TriggersHandler handlers the "/triggers" endpoint
func TriggersHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	x := "These are you're triggers"
	fmt.Fprintf(w, "Remedies: %v\n", x)
}

// RemediesCategoryHandler handles "/remedies{category}"
func RemediesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: %v\n", vars["category"])
}
