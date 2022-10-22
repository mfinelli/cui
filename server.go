// cui: http request/response tui
// Copyright 2022  Mario Finelli
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

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
