package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"sms/handlers"
	"sms/proto"
	grpc_from0 "sms/proto/github.com/monkrus/grpc-from0"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		fmt.Print("Cant' load env file")
	}
	go listenAndServeGrpc()
	listenAndServeHTTP()
}

func listenAndServeHTTP() {
	r := mux.NewRouter()
	r.HandleFunc("/generate/", handlers.PhoneNumberHandler).Methods("POST")
	r.HandleFunc("/verificate/", handlers.CodeVerificationHandler).Methods("POST")
	log.Print("HTTP server started...")
	log.Fatal(http.ListenAndServe(":80",r))
}

func listenAndServeGrpc() {
	var server proto.Server
	grpcServer := grpc.NewServer()
	grpc_from0.RegisterPayloadServer(grpcServer, server)
	listen, _ := net.Listen("tcp", ":3000")
	log.Print("GRPC server started...")
	log.Fatal(grpcServer.Serve(listen))
}
