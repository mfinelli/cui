package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

const devServerPort = ":7999"

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)["path"]

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "PATH: %s\n", path)
	fmt.Fprintf(w, "METHOD: %s\n\n", r.Method)

	fmt.Fprintf(w, "QUERY PARAMETERS:\n")
	for key, values := range r.URL.Query() {
		for _, value := range values {
			fmt.Fprintf(w, "%s: %s\n", key, value)
		}
	}
	fmt.Fprintf(w, "\n")
	// fmt.Fprintln(w, r.URL.Query())

	fmt.Fprintf(w, "HEADERS:\n")
	for header := range r.Header {
		for _, value := range r.Header[header] {
			fmt.Fprintf(w, "%s: %s\n", header, value)
		}
	}
	fmt.Fprintf(w, "\n")

	fmt.Fprintf(w, "BODY:\n")
	fmt.Fprintln(w, string(body))
}

func cuiDevServer(methods []string) {
	router := mux.NewRouter()
	router.HandleFunc("/", handlerFunc).Methods(methods...)
	router.HandleFunc("/{path}", handlerFunc).Methods(methods...)

	http.Handle("/", router)
	http.ListenAndServe(devServerPort, nil)
}
