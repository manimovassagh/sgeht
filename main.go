package sgeht

import (
	"io"
	"log"
	"net/http"
)

func SReq(url string) string {
	// godotenv packages

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(body))
	return string(body)

}
