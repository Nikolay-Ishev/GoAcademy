package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
)
func parseURL(rawurl string) (url url.URL) {
	urlPtr, err := url.Parse(rawurl)
	if err != nil {
		log.Printf("Error parsing URL: '%s': %v\n", rawurl, err)
	}
	return *urlPtr
}

func main() {
	file, err := os.Open("homework/usernames.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)


	scanner := bufio.NewScanner(file)

	var usernames []string

	for scanner.Scan() {
		// fmt.Println(scanner.Text())  // print the token in unicode-char
		usernames = append(usernames, scanner.Text())
	}

	fmt.Println(fmt.Sprintf("https://api.github.com/users/%v", usernames[0]))

	type Student struct {
		usersData   url.URL
		repData     url.URL
	}

	userInfo := []Student{
		{
			usersData: parseURL(fmt.Sprintf("https://api.github.com/users/%v", usernames[0])),
			repData:   parseURL(fmt.Sprintf("https://api.github.com/users/%v/repos", usernames[0])),
			//parseURL(fmt.Sprintf("https://api.github.com/repos/%v/${repo-name}/languages, usernames[0]")),
		},
	}

	data, err := json.Marshal(userInfo)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	// Prettier formatting
	data, err = json.MarshalIndent(userInfo, "", "     ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)




	var books []Student
	if err := json.Unmarshal(data, &books); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println("AFTER UNMARSHAL:\n")
	for i, book := range books {
		fmt.Printf("%d: %#v\n", i, book)
	}
}
