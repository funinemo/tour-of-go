package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount 文字列を分解してカウントする
// https://golang.org/pkg/strings/#Fields
// https://github.com/golang/tour/blob/master/wc/wc.go
func WordCount(s string) map[string]int {
	m := make(map[string]int)
	strArr := strings.Fields(s)
	for _, v := range strArr {
		m[v]++
	}
	//	return map[string]int{"Go!":1, "I":1, "am":1, "learning":1}
	return m
}

func main() {
	wc.Test(WordCount)
}
