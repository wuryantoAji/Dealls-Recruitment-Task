package main

import (
	"encoding/json"
	"fmt"
	"log"      // Package for logging
	"net/http" // Package for HTTP client and server
	"os"
	"os/signal"
	"signup-login/database"
	"signup-login/service"
)

// function for user sign up
func signUpNewUser(w http.ResponseWriter, r *http.Request) {
	err := service.UserSignUp(r.FormValue("name"), r.FormValue("username"), r.FormValue("password"))
	if err != nil {
		responseWriter(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWriter(w, http.StatusOK, "user sign up success, welcome "+r.FormValue("username"))
}

// function for user login
func loginUser(w http.ResponseWriter, r *http.Request) {
	err := service.UserLogin(r.FormValue("username"), r.FormValue("password"))
	if err != nil {
		responseWriter(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWriter(w, http.StatusOK, "user login success, hello "+r.FormValue("username"))
}

// function for writing a response to client
func responseWriter(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

// function that defines all available endpoint and make the program listen request
func handleRequests() {
	http.Handle("/signUp", http.HandlerFunc(signUpNewUser))
	http.Handle("/login", http.HandlerFunc(loginUser))
	fmt.Println("API ready to use")
	fmt.Println("Press ctrl+c to shutdown program")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// main function when running the program
func main() {
	// Setup a channel to receive a signal
	done := make(chan os.Signal, 1)

	// Notify this channel when a SIGINT is received
	signal.Notify(done, os.Interrupt)

	// Fire off a goroutine to loop until that channel receives a signal.
	// When a signal is received simply exit the program
	go func() {
		for _ = range done {
			fmt.Println("program is shutting down")
			database.CloseDB()
			os.Exit(0)
		}
	}()
	database.InitDB()
	handleRequests()
}
