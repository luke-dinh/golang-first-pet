package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/phuonglvh/pet/app/controller"
	"github.com/phuonglvh/pet/app/model"
	"github.com/phuonglvh/pet/util/logger"
)

func main() {
	logger.Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", controller.HomeHandler)
	myRouter.HandleFunc("/", controller.HomeHandler)
	myRouter.HandleFunc("/generator/", controller.ViewCodeHandler)
	myRouter.HandleFunc("/chat", controller.ChatIndex)
	myRouter.HandleFunc("/chat/rooms/{id}", controller.ChatIndex)

	chatHandler := &controller.ChatHandler{Rooms: make(map[string]*model.Room)}
	myRouter.Handle("/chat/rooms/{id}/ws", chatHandler)
	logger.Info.Println("Server is listening on port ", 8080)
	http.ListenAndServe(":8080", myRouter)
}
