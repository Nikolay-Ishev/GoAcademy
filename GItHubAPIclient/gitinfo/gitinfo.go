//Package gitinfo contains functions for processing information about a gitHub user
package gitinfo

import (
	"encoding/json"
	"fmt"
	"github.com/Nikolay-Ishev/Proxiad-GoAcademy/tree/master/GItHubAPIclient/datamanipulator"
	"github.com/Nikolay-Ishev/Proxiad-GoAcademy/tree/master/GItHubAPIclient/httprequest"
	"log"
)

//reads a gitHub username and returns information about the user
func GithubInfo(username string) (datamanipulator.GitInfo, map[string]int32) {
	gitUserUrl := fmt.Sprintf("https://api.github.com/users/%s", username)
	gitReposUrl := fmt.Sprintf("https://api.github.com/users/%s/repos", username)


	userResp := httprequest.HttpRequest(gitUserUrl)
	reposResp := httprequest.HttpRequest(gitReposUrl)

	var userInfo datamanipulator.GitInfo
	var reposInfo []datamanipulator.GitInfo

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
		langResp := httprequest.HttpRequest(gitLangUrl)
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