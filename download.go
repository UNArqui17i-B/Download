package main 

import (
		"github.com/ant0ine/go-json-rest/rest"
		"log"
		"net/http"
)

type RequestedFile struct{
	IdFile string
	UserEmail string
}

func ReceiveRequest(w rest.ResponseWriter, req *rest.Request) {
	m := req.URL.Query()

	fileRequest := RequestedFile{
		IdFile: 	req.PathParam("id"),
		UserEmail: 	m["email"][0],
	}
	w.WriteJson(&fileRequest)
}

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
			rest.Get("/download/:id", ReceiveRequest),
		)

	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":4025", api.MakeHandler()))
}