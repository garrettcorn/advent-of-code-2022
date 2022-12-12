package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func isDuplicateLetter(s string) bool {
	result := false
	for idx, val := range s {
		subString := s[:idx] + s[idx+1:]
		if strings.Contains(subString, string(val)) {
			result = true
		}
	}
	return result
}

func findStartOfPacketMarker(s string) int {
	for i := 4; i <= len(s); i++ {
		subString := s[i-4 : i]
		if !isDuplicateLetter(subString) {
			return i
		}
	}
	return 0
}

func findStartOfMessage(s string) int {
	for i := 14; i <= len(s); i++ {
		subString := s[i-14 : i]
		if !isDuplicateLetter(subString) {
			return i
		}
	}
	return 0
}

func getFileText(fp string) string {
	f, err := os.Open(fp)
	if err != nil {
		panic(err)
	}

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	return string(bs)
}

func main() {
	filePath := "./input"
	s := getFileText(filePath)
	i := findStartOfPacketMarker(s)
	fmt.Printf("PART1: start of packet marker is %v\n", i)
	fmt.Printf("PART2: start of message marker is %v\n", findStartOfMessage(s))
}
