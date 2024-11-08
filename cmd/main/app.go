package main

import (
	"awesomeProject/internal/user"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	log.Println("Create router")

	router := httprouter.New()

	log.Println("register logger handler")

	handler := user.NewHandler()

	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	log.Println("start server application")

	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("server is listening 127.0.0.1:1234")
	log.Fatalln(server.Serve(listener))
}
