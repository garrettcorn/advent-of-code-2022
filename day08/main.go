package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type tree struct {
	size int
}

type forest [][]tree

func (f forest) String() string {
	result := ""
	for x := 0; x < len(f); x++ {
		for y := 0; y < len(f[x]); y++ {
			result += fmt.Sprintf("%v", f[x][y].size)
		}
		result += "\n"
	}
	return result
}

func (f forest) NumVisible() int {
	count := 0
	for idy, val := range f {
		for idx, _ := range val {
			if f.IsVisible(idx, idy) {
				count++
			}
		}
	}
	return count
}

func reverse(s []tree) []tree {
	t := []tree{}
	for i := len(s) - 1; i >= 0; i-- {
		t = append(t, s[i])
	}
	return t
}

func (f forest) Left(x, y int) []tree {
	return reverse(f[y][:x])
}

func (f forest) Right(x, y int) []tree {
	return f[y][x+1:]
}

func (f forest) Above(x, y int) []tree {
	result := []tree{}
	for i := 0; i < y; i++ {
		result = append(result, f[i][x])
	}
	return reverse(result)
}

func (f forest) Below(x, y int) []tree {
	result := []tree{}
	for i := len(f) - 1; i > y; i-- {
		result = append(result, f[i][x])
	}
	return reverse(result)
}

func (f forest) VisibleLeft(x, y int) bool {
	// fmt.Println("f.Left(x, y)=", f.Left(x, y))
	for _, v := range f.Left(x, y) {
		if f[y][x].size <= v.size {
			return false
		}
	}
	return true
}

func (f forest) VisibleRight(x, y int) bool {
	// fmt.Println("f.Right(x, y)=", f.Right(x, y))
	for _, v := range f.Right(x, y) {
		if f[y][x].size <= v.size {
			return false
		}
	}
	return true
}

func (f forest) VisibleAbove(x, y int) bool {
	// fmt.Println("f.Above(x, y)=", f.Above(x, y))
	for _, v := range f.Above(x, y) {
		if f[y][x].size <= v.size {
			return false
		}
	}
	return true
}

func (f forest) VisibleBelow(x, y int) bool {
	// fmt.Println("f.Below(x, y)=", f.Below(x, y))
	for _, v := range f.Below(x, y) {
		if f[y][x].size <= v.size {
			return false
		}
	}
	return true
}

func (f forest) ScenicScoreLeft(x, y int) int {
	// fmt.Println("f.Left(x, y)=", f.Left(x, y))
	count := 0
	for _, v := range f.Left(x, y) {
		count++
		if f[y][x].size <= v.size {
			return count
		}
	}
	return count
}

func (f forest) ScenicScoreRight(x, y int) int {
	// fmt.Println("f.Right(x, y)=", f.Right(x, y))
	count := 0
	for _, v := range f.Right(x, y) {
		count++
		if f[y][x].size <= v.size {
			return count
		}
	}
	return count
}

func (f forest) ScenicScoreAbove(x, y int) int {
	// fmt.Println("f.Above(x, y)=", f.Above(x, y))
	count := 0
	for _, v := range f.Above(x, y) {
		count++
		if f[y][x].size <= v.size {
			return count
		}
	}
	return count
}

func (f forest) ScenicScoreBelow(x, y int) int {
	// fmt.Println("f.Below(x, y)=", f.Below(x, y))
	count := 0
	for _, v := range f.Below(x, y) {
		count++
		if f[y][x].size <= v.size {
			return count
		}
	}
	return count
}

func (f forest) highestScenicScore() int {
	highestScore := 0
	for idy, val := range f {
		for idx, _ := range val {
			score := f.ScenicScore(idx, idy)
			if score > highestScore {
				highestScore = score
			}
		}
	}
	return highestScore
}

func (f forest) ScenicScore(x, y int) int {
	// fmt.Printf("f.ScenicScoreLeft(%v, %v)=%v * f.ScenicScoreRight(x, y)=%v * f.ScenicScoreAbove(x, y)=%v * f.ScenicScoreBelow(x, y)=%v\n", x, y, f.ScenicScoreLeft(x, y), f.ScenicScoreRight(x, y), f.ScenicScoreAbove(x, y), f.ScenicScoreBelow(x, y))
	// fmt.Printf("ScenicScore = %v\n", f.ScenicScoreLeft(x, y)*f.ScenicScoreRight(x, y)*f.ScenicScoreAbove(x, y)*f.ScenicScoreBelow(x, y))
	return f.ScenicScoreLeft(x, y) * f.ScenicScoreRight(x, y) * f.ScenicScoreAbove(x, y) * f.ScenicScoreBelow(x, y)
}

func (f forest) IsVisible(x, y int) bool {
	if f.VisibleLeft(x, y) || f.VisibleRight(x, y) || f.VisibleAbove(x, y) || f.VisibleBelow(x, y) {
		return true
	}
	return false
}

func getForestFromInput(filePath string) forest {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	frst := forest{}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		row := []tree{}
		for _, val := range line {
			s, err := strconv.Atoi(string(val))
			if err != nil {
				panic(err)
			}
			row = append(row, tree{size: s})
		}
		frst = append(frst, row)
	}

	return frst
}

func main() {
	filePath := "./input"
	f := getForestFromInput(filePath)
	// fmt.Printf("%v\n", f)
	fmt.Printf("PART1: %v\n", f.NumVisible())
	fmt.Printf("PART2: %v\n", f.highestScenicScore())
}
