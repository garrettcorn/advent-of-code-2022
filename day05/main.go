package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type stack struct {
	crates []rune
}

func (s stack) String() string {
	var str string
	for _, val := range s.crates {
		str += fmt.Sprintf("[%v]\n", string(val))
	}
	return str
}

func (s *stack) AddCrate(r rune) {
	s.crates = append([]rune{r}, s.crates...)
}

func (s *stack) Push(cr []rune) {
	for _, val := range cr {
		s.AddCrate(val)
	}
}

func (s *stack) Pop(num int) []rune {
	r := s.crates[:num]
	s.crates = s.crates[num:]
	return r
}

func (s *stack) Push9001(cr []rune) {
	for i := len(cr); i > 0; i-- {
		s.AddCrate(cr[i-1])
	}
}

func (s *stack) Pop9001(num int) []rune {
	r := s.crates[:num]
	s.crates = s.crates[num:]
	return r
}

type cargo struct {
	stacks []stack
}

func (c cargo) TopOfEachStack() string {
	var s string
	for _, val := range c.stacks {
		if len(val.crates) >= 1 {
			s += string(val.crates[0])
		}
	}
	return s
}

func (c cargo) String() string {
	var s string
	for idx, val := range c.stacks {
		s += fmt.Sprintf("stack(%v)\n%v\n", idx+1, val)
	}
	return s
}

func (c cargo) Move(from, to, num int) {
	cr := c.stacks[from-1].Pop(num)
	c.stacks[to-1].Push(cr)
}

func (c cargo) Move9001(from, to, num int) {
	cr := c.stacks[from-1].Pop9001(num)
	c.stacks[to-1].Push9001(cr)
}

func (c cargo) Execute(p procedures) {
	for _, v := range p.moves {
		c.Move(v.fromStack, v.toStack, v.amount)
	}
}

func (c cargo) Execute9001(p procedures) {
	for _, v := range p.moves {
		c.Move9001(v.fromStack, v.toStack, v.amount)
	}
}

type move struct {
	amount    int
	fromStack int
	toStack   int
}

func (m move) String() string {
	return fmt.Sprintf("move %v from %v to %v", m.amount, m.fromStack, m.toStack)
}

type procedures struct {
	moves []move
}

func (p procedures) String() string {
	var s string
	for _, v := range p.moves {
		s += fmt.Sprintf("%v\n", v)
	}
	return s
}

func getCargoAndProcesures(filePath string) (cargo, procedures) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	c := cargo{stacks: []stack{}}
	p := procedures{moves: []move{}}
	firstLine := true
	cargoDone := false
	for fileScanner.Scan() {
		line := fileScanner.Text()
		stackNum := 0
		if !cargoDone {
			for i := 0; i < len(line)-2; i += 4 {
				r := rune(line[i+1])
				// fmt.Println(string(r))
				if string(r) == "1" {
					cargoDone = true
					break
				} else {

					if firstLine {
						s := stack{crates: []rune{}}
						c.stacks = append(c.stacks, s)
					}

					if string(r) != " " {
						c.stacks[stackNum].crates = append(c.stacks[stackNum].crates, r)
					}
					stackNum++
				}
			}
		}

		// if cargo is done it is time to start getting the procedures
		if cargoDone {
			if strings.HasPrefix(line, "move") {
				re := regexp.MustCompile("[0-9]+")
				ss := re.FindAllString(line, -1)
				if len(ss) >= 3 {
					a, err := strconv.Atoi(ss[0])
					if err != nil {
						panic(err)
					}
					f, err := strconv.Atoi(ss[1])
					if err != nil {
						panic(err)
					}
					t, err := strconv.Atoi(ss[2])
					if err != nil {
						panic(err)
					}
					p.moves = append(p.moves, move{amount: a, fromStack: f, toStack: t})
				}
			}
		}

		firstLine = false
	}

	return c, p
}

func main() {
	filePath := "./input"
	c, p := getCargoAndProcesures(filePath)
	c.Execute(p)
	log.Printf("PART1: TopOfEachStack = %v\n", c.TopOfEachStack())
	// log.Println(c)
	c, p = getCargoAndProcesures(filePath)
	c.Execute9001(p)
	log.Printf("PART2: TopOfEachStack = %v\n", c.TopOfEachStack())
	// log.Println(c)
}
