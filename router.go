package main

import(
	"github.com/gorilla/mux"
	"FileAccess/access"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/info/:id", access.GetInformation)
}
