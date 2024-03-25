package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func setupRoutes(router *mux.Router) {
	router.HandleFunc("/", RootHandler).Methods("GET")
}

func Serve(port int) {
	router := mux.NewRouter()
	setupRoutes(router)

	addr := fmt.Sprintf("localhost:%v", port)

	fmt.Printf("Listening on: %s\n", addr)

	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatalf("server encountered an error: %v", err)
	}
}
