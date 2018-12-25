package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWorld).Methods("GET")
	myRouter.HandleFunc("/_ah/health", healthCheckHandler).Methods("GET")
	myRouter.HandleFunc("/users", AllUsers).Methods("POST")
	myRouter.HandleFunc("/user/{name}/{email}", NewUser).Methods("POST")
	myRouter.HandleFunc("/user/{name}", DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", UpdateUser).Methods("PUT")
	log.Print("Listening on port :8081")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Better than good!")
}

func main() {
	handleRequests()
}
