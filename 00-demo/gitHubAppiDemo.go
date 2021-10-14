package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
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
func githubInfo(username string) (gitInfo, map[string]int32) {
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
	//contains info about distribution of programming languages
	allLang := make(map[string]int32)
	for i:=0; i<len(reposInfo); i++ {
		//reads the forks of each repo and adds the information in userInfo
		userInfo.Forks += reposInfo[i].Forks

		//reads the programming languages of each repo and adds the information in allLang
		repo := reposInfo[i].Name
		gitLangUrl := fmt.Sprintf("https://api.github.com/repos/%s/%s/languages", username, repo)
		langResp := httpRequest(gitLangUrl)
		langInfo := make(map[string]int32)
		if err := json.Unmarshal(langResp, &langInfo); err != nil {
			log.Fatalf("JSON unmarshaling failed: %s", err)
		}
		for k, v := range langInfo {
			allLang[k] += v
		}
	}
	return userInfo, allLang
}


// A data structure to hold a key/value pair.
type Pair struct {
	Key string
	Value int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

// A function to turn a map into a PairList, then sort and return it.
func sortMapByValue(m map[string]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	return p
}


// A data structure to hold gitHub user information
type gitInfo struct {
	Name                	string
	Location 				string
	Followers 				int32
	Public_repos			int32
	Forks               	int32
	Languages_distribution  string
}

//Turns the int values in a map to percent
func percentMap(m map[string]int32) map[string]int {
	var totalN float64
	for _, v := range m {
		totalN += float64(v)
	}

	percMap := make(map[string]int)
	for k, v := range m {
		perc := (float64(v) / totalN) * 100
		percMap[k] = int(perc)
	}
	return percMap
}


//Turns first 5 values of a PairList into a string
func formatPairList(p PairList) string {
	var langArr []string
	totalPerc := 0
	for i := range p {
		k := p[i].Key
		v := p[i].Value
		totalPerc += v
		s := fmt.Sprintf("%s-%v%%", k, v)
		langArr = append(langArr, s)
		if len(langArr) > 4 {
			otherLang := 100 - totalPerc
			x := fmt.Sprintf("Other languages-%v%%", otherLang)
			langArr = append(langArr, x)
			break
		}
	}
	str := strings.Join(langArr[:],", ")
	return str
}


func main() {
	fmt.Println("Please enter the file name below:")
	//var usernames []string
	usernames := "Nikolay-Ishev"
	//usernames = readFile()
	//for i:=0;i<len(usernames);i++ {
		//create files with the information returned from the requests
		//userInfo, userLang := githubInfo(usernames[i])
		userInfo, userLang := githubInfo(usernames)

		percLang := percentMap(userLang)

		sortedLang := sortMapByValue(percLang)

		userInfo.Languages_distribution = formatPairList(sortedLang)

		userData, err := json.MarshalIndent(userInfo, "", "     ")
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}

		fmt.Println(string(userData))
}


