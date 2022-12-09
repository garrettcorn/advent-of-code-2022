package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type elves []elf

func (es elves) Len() int {
	return len(es)
}

func (es elves) Less(i, j int) bool {
	return es[i].TotalCalories() > es[j].TotalCalories()
}

func (es elves) Swap(i, j int) {
	es[i], es[j] = es[j], es[i]
}

type elf struct {
	Calories []int
}

func (e elf) TotalCalories() int {
	sum := 0
	for _, value := range e.Calories {
		sum += value
	}
	return sum
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getElvesFromFile(filePath string) elves {
	readFile, err := os.Open(filePath)
	check(err)
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	es := elves{elf{Calories: []int{}}}
	i := 0

	for fileScanner.Scan() {
		num, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			es = append(es, elf{Calories: []int{}})
			i++
		} else {
			es[i].Calories = append(es[i].Calories, num)
		}

	}

	return es
}

func getElfWithHighestCalories(es elves) elf {
	elfWithHighestCalories := elf{Calories: []int{0}}
	for _, elf := range es {
		if elfWithHighestCalories.TotalCalories() < elf.TotalCalories() {
			elfWithHighestCalories = elf
		}
	}
	return elfWithHighestCalories
}

func getTopThreeElvesWithHighestCalories(es elves) elves {
	sort.Sort(es)
	if len(es) > 3 {
		return es[:3]
	} else {
		return es
	}
}

func sumElvesCalories(es elves) int {
	sum := 0
	for _, value := range es {
		sum += value.TotalCalories()
	}
	return sum
}

func main() {
	filePath := "./input"
	es := getElvesFromFile(filePath)
	elf := getElfWithHighestCalories(es)
	fmt.Printf("PART1: The elf with the highest calories has %v calories.\n", elf.TotalCalories())
	es = getTopThreeElvesWithHighestCalories(es)
	fmt.Printf("PART2: The top 3 elves have %v calories.\n", sumElvesCalories(es))
}
