package main 

import (
		"github.com/ant0ine/go-json-rest/rest"
		"log"
		"net/http"
		"fmt"
		"bytes"
		"encoding/json"
		//"github.com/leesper/couchdb-golang"
)

const DBurl string = "http://127.0.0.1:5984/blinkbox_files"

type RequestedFile struct{
	IdFile string
	UserEmail string
}

type FileInformation struct{
	Id string `json:"_id"`
	Name string `json:"name"`
	Extension string `json:"extension"`
	Size int `json:"size"`
	UploadedDate float64 `json:"uploaded_date"`
	ExpiringDate float64 `json:"expiring_date"`
	Owner string `json:"owner"`
	Shared []string `json:"shared"`
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

func ReceiveRequest(w rest.ResponseWriter, req *rest.Request) {
	m := req.URL.Query()

	fileRequest := RequestedFile{
		IdFile: 	req.PathParam("id"),
		UserEmail: 	m["email"][0],
	}

	GetAttachment(w, DBurl, fileRequest.IdFile, fileRequest.UserEmail)
}

func GetAttachment(w rest.ResponseWriter, db string, id string,  email string){
	var buffer bytes.Buffer

	buffer.WriteString(db)
	if db[len(db)-1:] != "/" {
		buffer.WriteString("/")
	}
	buffer.WriteString(id)
	url := buffer.String()

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Connection to DB response: ", err)
		return	
	}

	if resp.StatusCode != 200 {
		w.WriteHeader(resp.StatusCode)
		return
	}

	fmt.Printf("ID in DB result: %s\n", resp.Status)

	doc := new(FileInformation)
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&doc)
	fmt.Println(err)
	fmt.Println(doc)
	fmt.Println(doc.Shared)
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

	VerifyDatabaseExistance(DBurl)

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":4025", api.MakeHandler()))
}

func isValueInList(value string, list []string) bool{
	for _, curr := range list {
		if curr == value {
			return true
		}
	}
	return false
}