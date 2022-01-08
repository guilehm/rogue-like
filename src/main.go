package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Hello")

	r := mux.NewRouter()
	_ = http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
