package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type elf struct {
	from int
	to   int
}

func (e elf) FullyContainedBy(otherElf elf) bool {
	if e.from >= otherElf.from && e.to <= otherElf.to {
		return true
	}
	return false
}

func (e elf) OverlapWith(otherElf elf) bool {
	for i := e.from; i <= e.to; i++ {
		if i >= otherElf.from && i <= otherElf.to {
			return true
		}
	}
	return false
}

func (e elf) String() string {
	return fmt.Sprintf("%v-%v", e.from, e.to)
}

type elfPair struct {
	elfA elf
	elfB elf
}

func (ep elfPair) IsElfAContainedByElfB() bool {
	return ep.elfA.FullyContainedBy(ep.elfB)
}

func (ep elfPair) IsElfBContainedByElfA() bool {
	return ep.elfB.FullyContainedBy(ep.elfA)
}

func (ep elfPair) IsThereAFullyContainedElf() bool {
	if (ep.IsElfBContainedByElfA() == true) || (ep.IsElfAContainedByElfB() == true) {
		return true
	}
	return false
}

func (ep elfPair) Overlap() bool {
	return ep.elfA.OverlapWith(ep.elfB)
}

func (ep elfPair) String() string {
	return fmt.Sprintf("%v,%v", ep.elfA, ep.elfB)
}

type list []elfPair

func (l list) NumberOfOverlapPairs() int {
	num := 0
	for _, p := range l {
		if p.Overlap() {
			num++
		}
	}
	return num
}

func (l list) NumberOfFullyContainedPairs() int {
	num := 0
	for _, p := range l {
		if p.IsThereAFullyContainedElf() {
			num++
		}
	}
	return num
}

func (l list) String() string {
	var r string
	for _, p := range l {
		r += fmt.Sprintf("%v\n", p)
	}
	return r
}

func getListFromFile(filepath string) list {
	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	l := list{}

	for fileScanner.Scan() {
		text := fileScanner.Text()
		ss := strings.Split(text, ",")

		ePair := elfPair{}
		for idx, v := range ss {
			sss := strings.Split(v, "-")
			f, err := strconv.Atoi(sss[0])
			if err != nil {
				panic(err)
			}
			t, err := strconv.Atoi(sss[1])
			if err != nil {
				panic(err)
			}

			e := elf{from: f, to: t}
			if idx == 0 {
				ePair.elfA = e
			} else {
				ePair.elfB = e
			}
		}
		l = append(l, ePair)
	}

	return l
}

func main() {
	filePath := "./input"
	l := getListFromFile(filePath)
	// fmt.Println(l)
	fmt.Printf("PART1: Number of fully contained pairs is %v\n", l.NumberOfFullyContainedPairs())
	fmt.Printf("PART2: Number of pairs that overlap is %v\n", l.NumberOfOverlapPairs())
}
