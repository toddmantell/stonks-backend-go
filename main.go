package main

import (
	"fmt"
	"log"
	"net/http"
	s "strconv"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Printf("Running...")
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	value := GrahamFormula(2.1, 10.0, 2.3)
	// This is just to make this work, will need to figure out how to output all info
	fmt.Fprintf(w, "something %s", s.Itoa(int(value.highConservativeGrahamFormulaNumber)))
}
