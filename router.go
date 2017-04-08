package main

import(
	"github.com/gorilla/mux"
	"fileAccess/access"
	"net/http"
	"fmt"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/info", access.GetInformation)
	router.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Println("Default")
	})
	router.NotFoundHandler = http.HandlerFunc(notFound)
	fmt.Println("Serving")
	http.Handle("/", router)
	http.ListenAndServe(":4025", nil)
}
