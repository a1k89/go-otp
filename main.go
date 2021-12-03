package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
	"sms/handlers"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		fmt.Print("Cant' load env file")
	}
	r := mux.NewRouter()
	r.HandleFunc("/generate/", handlers.PhoneNumberHandler).Methods("POST")
	r.HandleFunc("/verificate/", handlers.CodeVerificationHandler).Methods("POST")
	err = http.ListenAndServe(":80", r)
	if err != nil {
		panic("Can't listen and serve :80")
	}
}
