package main

import (
	"fmt"
	"net/http"
)

/*
 * Validate Http request.
 */
func validateHttpRequest(path string, method string, r *http.Request, w http.ResponseWriter) bool {
	if r.URL.Path != path {
		http.Error(
			w,
			"404: File Not Found",
			http.StatusNotFound)
		return false
	}

	if r.Method != method {
		http.Error(
			w,
			"Method is not supported",
			http.StatusNotFound)

		return false
	}

	return true
}

/*
 * Function to handle Form web server call.
 * * is a pointer for request address
 */
func formHandler(w http.ResponseWriter, r *http.Request) {
	if !validateHttpRequest("/form", "POST", r, w) {
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name : %v\n Address : %v\n", name, address)
}

/*
 * Function to handle hello web server call.
 * * is a pointer for request address
 */
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if !validateHttpRequest("/hello", "GET", r, w) {
		return
	}

	fmt.Fprintf(w, "Hello!!")
}
