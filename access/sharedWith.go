package access

import(
	"net/http"
	"github.com/fjl/go-couchdb"
	"log"
	"github.com/gorilla/mux"
	"encoding/json"
	"os"
	"fmt"
)

func FilesSharedWith(rw http.ResponseWriter, req *http.Request)  {
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
		if isValueInList(email, element.Document.Shared) {
			result.Documents = append(result.Documents, element.Document)
		}
	}

	js, err := json.Marshal(result)
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(js)
	fmt.Println("Shared with request: 200")
}