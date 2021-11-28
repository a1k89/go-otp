package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"sms/handlers"
)

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/generate/", handlers.PhoneNumberHandler).Methods("POST")
	r.HandleFunc("/verificate/", handlers.CodeVerificationHandler).Methods("POST")
	err := http.ListenAndServe(":80", r)
	if err != nil {
		panic("Can't listen and serve :80")
	}
}
