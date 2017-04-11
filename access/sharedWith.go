package access

import(
	"net/http"
	"github.com/fjl/go-couchdb"
	"log"
	"github.com/gorilla/mux"
	"encoding/json"
)

func FilesSharedWith(rw http.ResponseWriter, req *http.Request)  {
	vars := mux.Vars(req)

	email := vars["email"]

	client, err := couchdb.NewClient(url, nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.DB(DBname)

	allDocuments := new(AllDocuments)

	db.AllDocs(allDocuments, couchdb.Options{"include_docs": true})
	result := new(SharedWith)

	for _, element := range allDocuments.Rows {
		if isValueInList(email, element.Document.Shared) {
			result.Documents = append(result.Documents, element.Document)
		}
	}

	js, err := json.Marshal(result)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(js)
}