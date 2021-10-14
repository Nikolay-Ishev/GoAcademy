package main

import (
	"encoding/json"
	"fmt"
	"github.com/Nikolay-Ishev/Proxiad-GoAcademy/tree/master/GItHubAPIclient/datamanipulator"
	"github.com/Nikolay-Ishev/Proxiad-GoAcademy/tree/master/GItHubAPIclient/gitinfo"
	"github.com/Nikolay-Ishev/Proxiad-GoAcademy/tree/master/GItHubAPIclient/txtreader"
	"log"
)


func main() {
	fmt.Println("Please enter the file name below:")
	var usernames []string
	usernames = txtreader.TxtReader()
	for i:=0;i<len(usernames);i++ {
		//create files with the information returned from the requests
		userInfo, userLang := gitinfo.GithubInfo(usernames[i])

		percLang := datamanipulator.PercentMap(userLang)

		sortedLang := datamanipulator.SortMapByValue(percLang)

		userInfo.Languages_distribution = datamanipulator.FormatPairList(sortedLang)

		userData, err := json.MarshalIndent(userInfo, "", "     ")
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}

		fmt.Println(string(userData))
	}
}
