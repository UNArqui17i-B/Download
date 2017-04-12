package access

import (
	"net/http"
	"bytes"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
	"os"
)

func GetInformation(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	fileID := vars["id"]

	if fileID != "" {
		var buffer bytes.Buffer

		// Get complete document URL
		buffer.WriteString(os.Getenv("Url") + "/" + os.Getenv("DBName") + "/")
		buffer.WriteString(fileID)
		url := buffer.String()

		resp, err := http.Get(url)
		if err != nil {
			log.Fatal("Connection to DB response: ", err)
			return
		}

		doc := new(FileInformation)
		defer resp.Body.Close()
		err = json.NewDecoder(resp.Body).Decode(&doc)
		if err != nil {
			log.Fatal(err)
		}

		js, err := json.Marshal(doc)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
		}

		if resp.StatusCode == http.StatusOK {
			rw.Header().Set("Content-Type", "application/json")
			rw.WriteHeader(http.StatusOK)
			rw.Write(js)
		} else {
			rw.WriteHeader(http.StatusNotFound)
		}
	}
}