package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type program struct {
	xi      []instruction
	name    string
	version string
}

func (p program) CRT(cycle int) string {
	s := ""
	i := 0
	registerX := 1
	for i < cycle {
		for _, inst := range p.xi {
			for z := 0; z < inst.cycles; z++ {
				// fmt.Printf("i=%v x=%v\n", i, registerX)
				if i%40 == registerX-1 || i%40 == registerX || i%40 == registerX+1 {
					s += "#"
				} else {
					s += "."
				}
				i++
				if i != 0 && i%40 == 0 {
					s += "\n"
				}
			}
			if i >= cycle {
				continue
			} else {
				registerX += inst.value
			}
		}
	}
	return s
}

func (p program) GetSignalStrengthAtCycle(cycle int) int {
	i := 0
	registerX := 1
	for i < cycle {
		for _, inst := range p.xi {
			i += inst.cycles
			if i >= cycle {
				continue
			} else {
				registerX += inst.value
			}
		}
	}
	return cycle * registerX
}

func (p program) String() string {
	s := fmt.Sprintf("Program name:%v\nProgram version:%v\n\n", p.name, p.version)
	for _, i := range p.xi {
		s += fmt.Sprintf("%v\n", i)
	}
	return s
}

type instruction struct {
	name   string
	value  int
	cycles int
}

func (i instruction) String() string {
	return fmt.Sprintf("%v\t%v\t(%v cycles)", i.name, i.value, i.cycles)
}

func getProgramFromFile(filePath string) program {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	p := program{name: filePath, version: "v0.0.0", xi: []instruction{}}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		xs := strings.Split(line, " ")
		switch xs[0] {
		case "addx":
			v, err := strconv.Atoi(xs[1])
			if err != nil {
				panic(err)
			}
			p.xi = append(p.xi, instruction{name: xs[0], value: v, cycles: 2})
		case "noop":
			p.xi = append(p.xi, instruction{name: xs[0], value: 0, cycles: 1})
		}
	}
	return p
}

func sumOfSignalStrengths(p program, cycles []int) int {
	sum := 0
	for _, c := range cycles {
		sum += p.GetSignalStrengthAtCycle(c)
	}
	return sum
}

func main() {
	filePath := "./input"
	p := getProgramFromFile(filePath)
	cycles := []int{20, 60, 100, 140, 180, 220}
	fmt.Printf("PART1: %v\n", sumOfSignalStrengths(p, cycles))

	cycle := 240
	fmt.Printf("PART2: \n%v\n", p.CRT(cycle))
}
