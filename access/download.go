package access

import (
		"github.com/ant0ine/go-json-rest/rest"
		"log"
		"net/http"
		"fmt"
		"bytes"
		"encoding/json"
		"github.com/fjl/go-couchdb"
)

type RequestedFile struct{
	IdFile string
	UserEmail string
}

/*type FileInformation struct{
	Id string `json:"_id"`
	Name string `json:"name"`
	Extension string `json:"extension"`
	Size int `json:"size"`
	UploadedDate float64 `json:"uploaded_date"`
	ExpiringDate float64 `json:"expiring_date"`
	Owner string `json:"owner"`
	Shared []string `json:"shared"`
}*/

type Result struct{
	Url string
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
	if err != nil {
		log.Fatal(err)
	}

	// To implement: Check if is the owner
	if isValueInList(email, doc.Shared) {
		c, err := couchdb.NewClient("http://127.0.0.1:5984/", nil)
		if err != nil {
			log.Fatal(err)
		}

		att, err := c.DB("blinkbox_files").Attachment(doc.Id, doc.Name + "." + doc.Extension, "")
		if err != nil {
			log.Fatal(err)
		}

		buffer.WriteString("/")
		buffer.WriteString(att.Name)

		result := Result{
			Url: buffer.String(),
		}

		w.WriteJson(&result)
		w.WriteHeader(http.StatusOK)
	}else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

/*func main() {
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
}*/

func isValueInList(value string, list []string) bool{
	for _, curr := range list {
		if curr == value {
			return true
		}
	}
	return false
}