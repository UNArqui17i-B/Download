package main 

import (
		"github.com/ant0ine/go-json-rest/rest"
		"log"
		"net/http"
		"fmt"
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

func VerifyDatabaseExistance(url string) {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		log.Fatal("Database connection request: ", err)
		return	
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Database connection do: ", err)
	}

	fmt.Printf("Database connection: %s\n", resp.Status)

	if resp.StatusCode == 404 {
		req, err = http.NewRequest("PUT", url, nil)
		if err != nil {
			log.Fatal("Database creation request: ", err)
			return
		}

		resp, err = client.Do(req)
		if err != nil {
			log.Fatal("Database creation do: ", err)
		}

		fmt.Printf("Database creation: %s\n", resp.Status)
	}
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

	VerifyDatabaseExistance("http://127.0.0.1:5984/blinkbox_files")

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":4025", api.MakeHandler()))
}