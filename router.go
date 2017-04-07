package main

import(
	"github.com/gorilla/mux"
	"fileAccess/access"
	"net/http"
	"fmt"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/info", access.GetInformation)
	router.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Println("Default")
	})
	fmt.Println("Serving")

	http.ListenAndServe(":4025", nil)
}
