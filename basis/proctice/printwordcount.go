package main

import (
	"fmt"
	"strings"
)

func printWordCount(sentence string) map[string]int {
	wordCountRes := map[string]int{}
	wordRes := strings.Split(sentence, " ")
	for _, v := range wordRes {
		val := wordCountRes[v]
		wordCountRes[v] = val + 1
	}
	return wordCountRes
}

func printWordCountTest() {
	sentence := "how do you do"
	fmt.Println(printWordCount(sentence))
}
