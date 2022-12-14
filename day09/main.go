package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type move struct {
	piece piece
	turn  int
}

type instruction struct {
	direction string
	amount    int
}

type piece struct {
	xLoc int
	yLoc int
	name string
}

type instructions []instruction

type board struct {
	pieces   []piece
	maxRange int
	turn     int
	history  []move
}

func (b *board) MovePiece(direction string, pidx int) {
	switch direction {
	case "R":
		b.pieces[pidx].xLoc++
	case "L":
		b.pieces[pidx].xLoc--
	case "U":
		b.pieces[pidx].yLoc++
	case "D":
		b.pieces[pidx].yLoc--
	}
}

func (b *board) MoveTailPiece(front, back int) {
	x := b.pieces[front].xLoc - b.pieces[back].xLoc
	y := b.pieces[front].yLoc - b.pieces[back].yLoc
	if x > b.maxRange {
		b.pieces[back].xLoc++
		if y >= b.maxRange {
			b.pieces[back].yLoc++
		} else if y <= (-1 * b.maxRange) {
			b.pieces[back].yLoc--
		}
	} else if x < (-1 * b.maxRange) {
		b.pieces[back].xLoc--
		if y >= b.maxRange {
			b.pieces[back].yLoc++
		} else if y <= (-1 * b.maxRange) {
			b.pieces[back].yLoc--
		}
	} else if y > b.maxRange {
		b.pieces[back].yLoc++
		if x >= b.maxRange {
			b.pieces[back].xLoc++
		} else if x <= (-1 * b.maxRange) {
			b.pieces[back].xLoc--
		}
	} else if y < (-1 * b.maxRange) {
		b.pieces[back].yLoc--
		if x >= b.maxRange {
			b.pieces[back].xLoc++
		} else if x <= (-1 * b.maxRange) {
			b.pieces[back].xLoc--
		}
	}
}

func (b *board) RecordHistory() {
	for _, val := range b.pieces {
		b.history = append(b.history, move{piece: val, turn: b.turn})
	}
}

func (b *board) ExecuteInstruction(i instruction) {
	for a := 0; a < i.amount; a++ {
		b.turn++
		b.MovePiece(i.direction, 0)
		for idx, _ := range b.pieces {
			if idx != 0 {
				b.MoveTailPiece(idx-1, idx)
			}
		}
		b.RecordHistory()
		// fmt.Println(b.PrintBoard(6, 5))
	}
}

func (b *board) PrintBoard(width, height int) (s string) {
	for h := height - 1; h >= 0; h-- {
		for w := 0; w < width; w++ {
			value := "-"
			for _, p := range b.pieces {
				if w == p.xLoc && h == p.yLoc {
					if value == "-" {
						value = p.name
					} else {
						value = "M"
					}

				}
			}
			s += value
		}
		s += "\n"
	}
	return s
}

func (b *board) ExecuteInstructions(is instructions) {
	for _, v := range is {
		b.ExecuteInstruction(v)
	}
}

func (b *board) NumVisits(pIdx int) int {
	m := map[piece]bool{}
	for _, val := range b.history {
		if val.piece.name == b.pieces[pIdx].name {
			m[val.piece] = true
		}
	}
	return len(m)
}

func NewBoard(maxRange, numberOfTailPieces int) board {
	headPiece := piece{
		xLoc: 0,
		yLoc: 0,
		name: "H",
	}
	pieces := []piece{headPiece}
	history := []move{{piece: headPiece, turn: 0}}
	for i := 1; i <= numberOfTailPieces; i++ {
		p := piece{
			xLoc: 0,
			yLoc: 0,
			name: fmt.Sprintf("%v", i),
		}
		pieces = append(pieces, p)
		history = append(history, move{piece: p, turn: 0})
	}

	return board{
		pieces:   pieces,
		maxRange: maxRange,
		turn:     0,
		history:  history,
	}
}

func getInstructionFromFile(filePath string) (inst instructions) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		text := fileScanner.Text()
		xs := strings.Split(text, " ")
		a, err := strconv.Atoi(xs[1])
		if err != nil {
			panic(err)
		}
		inst = append(inst, instruction{direction: xs[0], amount: a})
	}
	return inst
}

func main() {
	filePath := "./input"
	instr := getInstructionFromFile(filePath)
	maxRange := 1
	numTailPieces := 1
	b := NewBoard(maxRange, numTailPieces)
	b.ExecuteInstructions(instr)
	fmt.Printf("PART1: %v\n", b.NumVisits(1))

	maxRange = 1
	numTailPieces = 9
	b = NewBoard(maxRange, numTailPieces)
	b.ExecuteInstructions(instr)
	fmt.Printf("PART2: %v\n", b.NumVisits(9))

}
