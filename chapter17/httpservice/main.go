package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8111
	http.HandleFunc("/", systemdServer)
	log.Printf("Server running - port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func systemdServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello !!! you are getting response from app running via systemd\n")
}
