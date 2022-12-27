package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

type item struct {
	worry big.Int
}

func (i item) String() string {
	return fmt.Sprintf("%v", i.worry.Text(10))
}

type throwTo struct {
	t int
	f int
}

type monkey struct {
	items       []item
	opperation  func(*big.Int)
	test        func(big.Int) bool
	action      throwTo
	inspections int
}

func (m monkey) String() (s string) {
	s += "items: "
	for _, i := range m.items {
		s += fmt.Sprintf("%v ", i)
	}
	s = strings.TrimSpace(s)
	return s
}

func (m *monkey) Inspect(itemIndex int) {
	m.opperation(&m.items[itemIndex].worry)
}

func (m *monkey) AddItem(i item) {
	m.items = append(m.items, i)
}

func (m *monkey) RemoveItem(itemIndex int) {
	m.items = append(m.items[:itemIndex], m.items[itemIndex+1:]...)
}

func (m *monkey) ModifyItemWorry(itemIndex int, f func(*big.Int)) {
	f(&m.items[itemIndex].worry)
}

func (m *monkey) Test(itemIndex int) bool {
	return m.test(m.items[itemIndex].worry)
}

func (m *monkey) ThrowItemTo(itemIndex int) int {
	if m.test(m.items[itemIndex].worry) {
		return m.action.t
	}
	return m.action.f
}

type monkeyInTheMiddle struct {
	monkeys []monkey
}

func (mitm monkeyInTheMiddle) String() (s string) {
	for idx, m := range mitm.monkeys {
		s += fmt.Sprintf("Monkey %v\n%v\n", idx, m)
	}
	return s
}

func getMitm(filePath string) monkeyInTheMiddle {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	mitm := monkeyInTheMiddle{monkeys: []monkey{}}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.HasPrefix(line, "Monkey") {
			m := monkey{items: []item{}, action: throwTo{}, inspections: 0}
			monkeyDone := false
			for fileScanner.Scan() && !monkeyDone {
				line := fileScanner.Text()
				xs := strings.Split(strings.TrimSpace(line), " ")
				switch strings.ToLower(xs[0]) {
				case "starting":
					for i := 2; i < len(xs); i++ {
						valString := strings.TrimSuffix(xs[i], ",")
						val, err := strconv.Atoi(valString)
						if err != nil {
							panic(err)
						}
						m.items = append(m.items, item{worry: *big.NewInt(int64(val))})
					}
				case "operation:":
					op := xs[4]
					secondVar := xs[5]
					if op == "+" {
						if secondVar == "old" {
							m.opperation = func(old *big.Int) { old.Sub(old, old) }
						} else {
							i, err := strconv.Atoi(secondVar)
							if err != nil {
								panic(err)
							}
							m.opperation = func(old *big.Int) { old.Add(old, big.NewInt(int64(i))) }
						}
					} else if op == "*" {
						if secondVar == "old" {
							m.opperation = func(old *big.Int) { old.Mul(old, old) }
						} else {
							i, err := strconv.Atoi(secondVar)
							if err != nil {
								panic(err)
							}
							m.opperation = func(old *big.Int) { old.Mul(old, big.NewInt(int64(i))) }
						}
					}

				case "test:":
					i, err := strconv.Atoi(xs[3])
					if err != nil {
						panic(err)
					}
					m.test = func(itemWorry big.Int) bool {
						return big.NewInt(0).Mod(&itemWorry, big.NewInt(int64(i))).Cmp(big.NewInt(0)) == 0
					}
				case "if":
					i, err := strconv.Atoi(xs[5])
					if err != nil {
						panic(err)
					}
					if strings.ToLower(xs[1]) == "true:" {
						m.action.t = i
					} else if strings.ToLower(xs[1]) == "false:" {
						m.action.f = i
						mitm.monkeys = append(mitm.monkeys, m)
						monkeyDone = true
					}
				}
			}
		}
	}
	return mitm
}

func (mitm *monkeyInTheMiddle) Round(rounds int, manageWorry bool) {
	for i := 0; i < rounds; i++ {
		fmt.Printf("i: %v\n", i)
		for idx := range mitm.monkeys {
			for range mitm.monkeys[idx].items {
				mitm.monkeys[idx].Inspect(0)
				if manageWorry {
					mitm.monkeys[idx].ModifyItemWorry(0, func(i *big.Int) { i.Div(i, big.NewInt(3)) })
				}
				throwItemto := mitm.monkeys[idx].ThrowItemTo(0)
				mitm.monkeys[throwItemto].AddItem(mitm.monkeys[idx].items[0])
				mitm.monkeys[idx].RemoveItem(0)
				mitm.monkeys[idx].inspections++
			}
		}
	}
}

func (mitm monkeyInTheMiddle) MonkeyBusiness() int {
	inspections := []int{}
	for _, m := range mitm.monkeys {
		inspections = append(inspections, m.inspections)
	}
	sort.Ints(inspections)
	one := inspections[len(inspections)-1]
	two := inspections[len(inspections)-2]
	return one * two
}

func (mitm monkeyInTheMiddle) StringAllMonkeyBusiness() (s string) {
	for idx, m := range mitm.monkeys {
		s += fmt.Sprintf("Monkey %v inspected items %v times.\n", idx, m.inspections)
	}
	for idx, m := range mitm.monkeys {
		s += fmt.Sprintf("Monkey %v\n%v\n", idx, m)
	}
	return s
}

func main() {
	filePath := "./input"
	mitm := getMitm(filePath)
	rounds := 20
	manageWorry := true
	mitm.Round(rounds, manageWorry)
	// fmt.Printf("%v", mitm.StringAllMonkeyBusiness())
	fmt.Printf("PART1: \n%v\t(sample=10605)\n\n", mitm.MonkeyBusiness())

	mitm = getMitm(filePath)
	rounds = 1000
	manageWorry = false
	mitm.Round(rounds, manageWorry)
	fmt.Printf("%v", mitm.StringAllMonkeyBusiness())
	fmt.Printf("PART2: \n%v\n\n", mitm.MonkeyBusiness())
}
