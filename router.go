package main

import (
	"github.com/gorilla/mux"
	"fileAccess/access"
	"net/http"
	"os"
)

var DefaultValues = map[string]string{
	"DB_PORT": "5984",
	"DB_URL": "127.0.0.1",
	"DB_NAME": "blinkbox_files",
	"HOST_PORT": "4025",
	"HOST_URL": "0.0.0.0"}

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

	os.Setenv("DB_URL", "http://" + os.Getenv("DB_URL") + ":" + os.Getenv("DB_PORT"))

	access.VerifyDatabaseExistance(os.Getenv("DB_URL") + "/" + os.Getenv("DB_NAME"))

	http.Handle("/", router)
	http.ListenAndServe(os.Getenv("HOST_URL") + ":" + os.Getenv("HOST_PORT"), nil)
}
