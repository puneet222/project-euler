package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func getScore(name string) int {
	name = strings.ReplaceAll(name, "\"", "")
	score := 0
	for _, v := range name {
		score += int(v) - 64
	}
	return score
}

func main() {
	data, err := ioutil.ReadFile("names.txt")
	if err != nil {
		log.Fatalln(err)
	}
	names := strings.Split(string(data), ",")
	sort.Strings(names)
	nameScores := 0
	for i, name := range names {
		score := getScore(name)
		nameScores += (i+1)*score
	}
	fmt.Printf("Total Name Scores: %v\n", nameScores)
}
