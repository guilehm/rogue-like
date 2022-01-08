package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello")

	r := mux.NewRouter()
	_ = http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
