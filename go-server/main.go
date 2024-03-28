package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = 8080

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at PORT %v\n", PORT)

	port_path := fmt.Sprintf(":%v", PORT)
	if err := http.ListenAndServe(port_path, nil); err != nil {
		log.Fatal(err)
	}
}
