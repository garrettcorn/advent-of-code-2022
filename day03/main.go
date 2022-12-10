package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getPriority(c rune) int {
	switch c {
	case 'a':
		return 1
	case 'b':
		return 2
	case 'c':
		return 3
	case 'd':
		return 4
	case 'e':
		return 5
	case 'f':
		return 6
	case 'g':
		return 7
	case 'h':
		return 8
	case 'i':
		return 9
	case 'j':
		return 10
	case 'k':
		return 11
	case 'l':
		return 12
	case 'm':
		return 13
	case 'n':
		return 14
	case 'o':
		return 15
	case 'p':
		return 16
	case 'q':
		return 17
	case 'r':
		return 18
	case 's':
		return 19
	case 't':
		return 20
	case 'u':
		return 21
	case 'v':
		return 22
	case 'w':
		return 23
	case 'x':
		return 24
	case 'y':
		return 25
	case 'z':
		return 26
	case 'A':
		return 27
	case 'B':
		return 28
	case 'C':
		return 29
	case 'D':
		return 30
	case 'E':
		return 31
	case 'F':
		return 32
	case 'G':
		return 33
	case 'H':
		return 34
	case 'I':
		return 35
	case 'J':
		return 36
	case 'K':
		return 37
	case 'L':
		return 38
	case 'M':
		return 39
	case 'N':
		return 40
	case 'O':
		return 41
	case 'P':
		return 42
	case 'Q':
		return 43
	case 'R':
		return 44
	case 'S':
		return 45
	case 'T':
		return 46
	case 'U':
		return 47
	case 'V':
		return 48
	case 'W':
		return 49
	case 'X':
		return 50
	case 'Y':
		return 51
	case 'Z':
		return 52
	}
	return 0
}

func getPriorities(c string) int {
	sum := 0
	for _, v := range c {
		sum += getPriority(v)
	}
	return sum
}

type compartment string

type rucksack struct {
	CompartmentA compartment
	CompartmentB compartment
}

type list struct {
	Rucksacks []rucksack
}

func getListFromFile(filePath string) list {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	l := list{Rucksacks: []rucksack{}}

	for fileScanner.Scan() {
		text := fileScanner.Text()
		// fmt.Printf("text='%v', len(text)=%v\n", text, len(text))
		cmp1 := text[:len(text)/2]
		cmp2 := text[len(text)/2:]

		// fmt.Printf("cmp1='%v' len(cmp1)=%v\n", cmp1, len(cmp1))
		// fmt.Printf("cmp2='%v' len(cmp2)=%v\n", cmp2, len(cmp2))

		l.Rucksacks = append(l.Rucksacks, rucksack{CompartmentA: compartment(cmp1), CompartmentB: compartment(cmp2)})
	}

	return l
}

func (r rucksack) GetEntireRucksack() string {
	return fmt.Sprintf("%v%v", r.CompartmentA, r.CompartmentB)
}

func (r rucksack) InRucksack(s string) bool {
	return strings.Contains(r.GetEntireRucksack(), s)
}

func (r rucksack) InBothCompartments() string {
	inBoth := ""
	for _, v := range r.CompartmentA {
		if !strings.Contains(inBoth, string(v)) && strings.Contains(string(r.CompartmentB), string(v)) {
			inBoth += string(v)
		}
	}
	return inBoth
}

func (r rucksack) ReorganizePriority() int {
	inBoth := r.InBothCompartments()
	fmt.Printf("inBoth = '%v' prioritiesValue=%v\n", inBoth, getPriorities(inBoth))
	return getPriorities(inBoth)
}

func (l list) SumOfPriorities() int {
	sum := 0
	for _, r := range l.Rucksacks {
		sum += r.ReorganizePriority()
	}
	return sum
}

type groups []group

type group struct {
	Rucksacks []rucksack
}

func (g groups) GetBadges() []rune {
	var runes []rune
	for _, v := range g {
		runes = append(runes, v.GetBadge())
	}
	return runes
}

func (g group) GetBadge() rune {
	fsack := ""
	if len(g.Rucksacks) >= 1 {
		fsack = g.Rucksacks[0].GetEntireRucksack()
	}

	for _, r := range fsack {
		for _, s := range g.Rucksacks {
			if !s.InRucksack(string(r)) {
				fsack = strings.ReplaceAll(fsack, string(r), "")
			}
		}
	}
	return rune(fsack[0])
}

func getGroupsFromList(l list) groups {
	g := groups{}
	groupSize := 3
	for idx, _ := range l.Rucksacks {
		if idx%groupSize == groupSize-1 {
			gr := group{Rucksacks: []rucksack{}}
			for i := groupSize - 1; i >= 0; i-- {
				gr.Rucksacks = append(gr.Rucksacks, l.Rucksacks[idx-i])
			}
			g = append(g, gr)
		} else if idx == len(l.Rucksacks) {
			g = append(g)
		}
	}
	return g
}

func main() {
	filePath := "./input"
	l := getListFromFile(filePath)
	fmt.Printf("PART1: Sum of all priorities is %v\n", l.SumOfPriorities())
	// l = list{Rucksacks: l.Rucksacks[:9]}
	g := getGroupsFromList(l)
	badges := string(g.GetBadges())
	fmt.Printf("badges=%v\n", badges)
	fmt.Printf("PART2: badges priority sum = %v\n", getPriorities(badges))
}
