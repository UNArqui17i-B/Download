package main

import (
	"github.com/gorilla/mux"
	"fileAccess/access"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/info/{id}", access.GetInformation)
	router.HandleFunc("/download/{id}/{email}", access.Download)
	router.HandleFunc("/sharedWith/{email}", access.FilesSharedWith)

	router.NotFoundHandler = http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusNotFound)
	})

	access.VerifyDatabaseExistance(access.DBurl)

	http.Handle("/", router)
	http.ListenAndServe(":4025", nil)
}
