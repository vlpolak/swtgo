package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, v := range strings.Split(s, " ") {
		fmt.Println(v)
		m[v]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
