package access

import(
	"net/http"
	"log"
	"fmt"
)

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

func isValueInList(value string, list []string) bool {
	for _, curr := range list {
		if curr == value {
			return true
		}
	}
	return false
}