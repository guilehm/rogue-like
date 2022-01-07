package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	_ = http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
