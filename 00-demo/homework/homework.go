package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

//reads a string with http request and returns the response in a byte array
func httpRequest(urlStr string) []byte {
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
	fmt.Println("Response status:", resp.Status)

	//read data from the response body and returns it in a byte array
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return b
}


//reads a file from the console and returns a string array with the information
func readFile() []string {
	var f string
	_, err := fmt.Scanln(&f)
	if err != nil {
		log.Fatalf("wrong input")
	}

	//Paths should start after the project directory
	file, err := os.Open(f)
	if err != nil {
		log.Fatalf("wrong input, please enter the path after the project directory")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	var usernames []string
	for scanner.Scan() {
		// fmt.Println(scanner.Text())  // print the token in unicode-char
		usernames = append(usernames, scanner.Text())
	}
	return usernames
}

//reads a github username and returns information about the user
func githubInfo(username string) gitInfo {
	gitUserUrl := fmt.Sprintf("https://api.github.com/users/%s", username)
	gitReposUrl := fmt.Sprintf("https://api.github.com/users/%s/repos", username)

	userResp := httpRequest(gitUserUrl)
	reposResp := httpRequest(gitReposUrl)

	var userInfo gitInfo
	var reposInfo []gitInfo

	if err := json.Unmarshal(userResp, &userInfo); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}

	if err := json.Unmarshal(reposResp, &reposInfo); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}

	for i:=0; i<len(reposInfo); i++ {
		userInfo.Forks += reposInfo[i].Forks
		userInfo.LanguagesArr = append(userInfo.LanguagesArr, reposInfo[i].Language)
	}
	return userInfo
}


type gitInfo struct {
	Name                string
	Location 			string
	Followers 			int32
	Public_repos		int32
	Forks				int32
	Language			string
	LanguagesArr 		[]string
}


func main() {
	fmt.Println("Please enter the file name below:")
	var usernames []string
	usernames = readFile()

	//create files with the information returned from the requests
	userInfo := githubInfo(usernames[0])

	fmt.Println(userInfo)
}






//-----------------------------------------------------------------------------------------------------//

//data, err := json.Marshal(userInfo)
//if err != nil {
//	log.Fatalf("JSON marshaling failed: %s", err)
//}
//fmt.Printf("%s\n", data)
//
//// Prettier formatting
//data, err = json.MarshalIndent(userInfo, "", "     ")
//if err != nil {
//	log.Fatalf("JSON marshaling failed: %s", err)
//}
//fmt.Printf("%s\n", data)
//
//
//
//
//var books []Student
//if err := json.Unmarshal(data, &books); err != nil {
//	log.Fatalf("JSON unmarshaling failed: %s", err)
//}
//fmt.Println("AFTER UNMARSHAL:\n")
//for i, book := range books {
//	fmt.Printf("%d: %#v\n", i, book)
//}

//type Student struct {
//	usersData   url.URL
//	repData     url.URL
//}
//
//userInfo := []Student{
//	{
//		usersData: parseURL(fmt.Sprintf("https://api.github.com/users/%v", usernames[0])),
//		repData:   parseURL(fmt.Sprintf("https://api.github.com/users/%v/repos", usernames[0])),
//		//parseURL(fmt.Sprintf("https://api.github.com/repos/%v/${repo-name}/languages, usernames[0]")),
//	},
//}

//func parseURL(rawUrl string) (url url.URL) {
	//	urlPtr, err := url.Parse(rawUrl)
	//	if err != nil {
	//		log.Printf("Error parsing URL: '%s': %v\n", rawUrl, err)
	//	}
	//	return *urlPtr
	//}