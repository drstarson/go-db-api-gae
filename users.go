package main

import (
	"fmt"
	"net/http"
)

func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All Users")
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New User")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update")
}
