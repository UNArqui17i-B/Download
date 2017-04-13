package main

import (
	"github.com/gorilla/mux"
	"fileAccess/access"
	"net/http"
	"os"
)

var DefaultValues = map[string]string{
	"Url": "http://127.0.0.1:5984",
	"DBName": "blinkbox_files"}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/fileAccess/info/{id}", access.GetInformation)
	router.HandleFunc("/fileAccess/download/{id}/{email}", access.Download)
	router.HandleFunc("/fileAccess/sharedWith/{email}", access.FilesSharedWith)
	router.HandleFunc("/fileAccess/ownedBy/{email}", access.OwnedBy)

	router.NotFoundHandler = http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusNotFound)
	})

	for key, val := range DefaultValues {
		if len(os.Getenv(key)) == 0 {
			os.Setenv(key, val)
		}
	}

	access.VerifyDatabaseExistance(os.Getenv("Url") + "/" + os.Getenv("DBName"))

	http.Handle("/", router)
	http.ListenAndServe(":4025", nil)
}
