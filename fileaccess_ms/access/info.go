package access

import (
	//"fmt"
	"net/http"
	"bytes"
	"log"
	"encoding/json"
)


func GetInformation(rw http.ResponseWriter, req *http.Request) {
	VerifyDatabaseExistance(DBurl);

	fileID := req.URL.Query().Get("id")

	if fileID != "" {
		var buffer bytes.Buffer

		// Get complete document URL
		buffer.WriteString(DBurl)
		buffer.WriteString("/")
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

		json.NewEncoder(rw).Encode(&doc)
		rw.WriteHeader(http.StatusOK)
	}
}