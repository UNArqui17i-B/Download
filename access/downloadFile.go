package access

import(
	"github.com/gorilla/mux"
	"net/http"
	"github.com/fjl/go-couchdb"
	"log"
	"io"
	"os"
)

func Download(rw http.ResponseWriter, req *http.Request)  {
	vars := mux.Vars(req)

	fileID := vars["id"]
	email := vars["email"]

	client, err := couchdb.NewClient(os.Getenv("DB_URL"), nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.DB(os.Getenv("DB_NAME"))

	doc := new(FileInformation)
	err = db.Get(fileID, &doc, nil)

	if email == doc.Owner || isValueInList(email, doc.Shared){
		att, err := db.Attachment(fileID, doc.Name + "." + doc.Extension, "")
		if err != nil{
			log.Fatal(err)
		}

		_, err = io.Copy(rw, att.Body)
		if err != nil {
			log.Fatal(err)
		}
	}else {
		rw.WriteHeader(http.StatusUnauthorized)
	}
}