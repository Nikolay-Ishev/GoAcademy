// Package datamanipulator contains functions for manipulating data structures
package datamanipulator

import (
	"fmt"
	"sort"
	"strings"
)

// A data structure to hold a key/value pair.
type Pair struct {
	Key string
	Value int
}

// A data structure to hold gitHub user information
type GitInfo struct {
	Name                	string
	Location 				string
	Followers 				int32
	Public_repos			int32
	Forks               	int32
	Languages_distribution  string
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

// A function to turn a map into a PairList, then sort and return it.
func SortMapByValue(m map[string]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	return p
}

//Converts int numbers of a map into percent
func PercentMap(m map[string]int32) map[string]int {
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

//Returns a string with the first 5 values of a PairList
func FormatPairList(p PairList) string {
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
