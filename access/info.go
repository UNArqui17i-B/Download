package access

import (
	"fmt"
	"net/http"
)

func GetInformation(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("Get information function")
}