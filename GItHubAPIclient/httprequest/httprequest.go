// Package httprequest contains function for processing http request
package httprequest

import (
	"io"
	"log"
	"net/http"
)

//reads a string with http request and returns the response in a byte array
func HttpRequest(urlStr string) []byte {
	//read string and returns request
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		log.Fatal(err)
	}

	//sends req and return resp
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	//fmt.Println("Response status:", resp.Status)

	//read data from the response body and returns it in a byte array
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return b
}
