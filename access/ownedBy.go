package access

import (
	"net/http"
	"github.com/gorilla/mux"
	"os"
	"log"
	"encoding/json"
	"github.com/fjl/go-couchdb"
)

func OwnedBy(rw http.ResponseWriter, req *http.Request)  {
	vars := mux.Vars(req)
	email := vars["email"]

	client, err := couchdb.NewClient(os.Getenv("DB_URL"), nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.DB(os.Getenv("DB_NAME"))

	allDocuments := new(AllDocuments)

	db.AllDocs(allDocuments, couchdb.Options{"include_docs": true})
	result := new(SharedWith)

	for _, element := range allDocuments.Rows {
		if email == element.Document.Owner {
			result.Documents = append(result.Documents, element.Document)
		}
	}

	js, err := json.Marshal(result)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(js)
}
