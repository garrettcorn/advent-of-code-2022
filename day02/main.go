package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type round struct {
	OpponentPlay play
	MyPlay       play
}

type play struct {
	Str   string
	Value int
	Orig  string
}

func toPlay(s string) play {
	p := play{}
	if (s == "A") || (s == "X") || (strings.ToLower(s) == "rock") {
		p = play{
			Str:   "Rock",
			Value: 1,
			Orig:  s,
		}
	} else if (s == "B") || (s == "Y") || (strings.ToLower(s) == "paper") {
		p = play{
			Str:   "Paper",
			Value: 2,
			Orig:  s,
		}
	} else if (s == "C") || (s == "Z") || (strings.ToLower(s) == "scissors") {
		p = play{
			Str:   "Scissors",
			Value: 3,
			Orig:  s,
		}
	}
	return p
}

func (r round) Draw() bool {
	if r.MyPlay.Str == r.OpponentPlay.Str {
		return true
	}
	return false
}

func (r round) Win() bool {
	win := false
	if (r.OpponentPlay.Str == "Rock") && (r.MyPlay.Str == "Paper") {
		win = true
	} else if (r.OpponentPlay.Str == "Paper") && (r.MyPlay.Str == "Scissors") {
		win = true
	} else if (r.OpponentPlay.Str == "Scissors") && (r.MyPlay.Str == "Rock") {
		win = true
	}
	return win
}

func (r round) Score() int {
	score := 0

	// Score for MyPlay
	score += r.MyPlay.Value

	// Score for outcome
	// Draw (3 points)
	if r.Draw() {
		score += 3
	} else if r.Win() {
		score += 6
	}
	return score
}

type game []round

func (g game) Score() int {
	sum := 0
	for _, round := range g {
		sum += round.Score()
	}
	return sum
}

func getGameFromFile(filePath string) game {
	readFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	g := game{}

	for fileScanner.Scan() {
		g = append(g, getRound(fileScanner.Text()))
	}

	return g
}

func getRound(text string) round {
	s := strings.Split(text, " ")
	return round{
		OpponentPlay: toPlay(s[0]),
		MyPlay:       toPlay(s[1]),
	}
}

func (g game) GameScore() int {
	score := 0
	for _, round := range g {
		score += round.Score()
	}
	return score
}

func (r *round) UpdateRoundToOutcome() {
	if r.MyPlay.Orig == "X" { // lose
		if r.OpponentPlay.Str == "Rock" {
			r.MyPlay = toPlay("Scissors")
		} else if r.OpponentPlay.Str == "Paper" {
			r.MyPlay = toPlay("Rock")
		} else if r.OpponentPlay.Str == "Scissors" {
			r.MyPlay = toPlay("Paper")
		}
	} else if r.MyPlay.Orig == "Y" { // draw
		r.MyPlay = r.OpponentPlay
	} else if r.MyPlay.Orig == "Z" { // win
		if r.OpponentPlay.Str == "Rock" {
			r.MyPlay = toPlay("Paper")
		} else if r.OpponentPlay.Str == "Paper" {
			r.MyPlay = toPlay("Scissors")
		} else if r.OpponentPlay.Str == "Scissors" {
			r.MyPlay = toPlay("Rock")
		}
	}
}

func (g *game) UpdateGameToOutcome() {
	for idx, _ := range *g {
		(*g)[idx].UpdateRoundToOutcome()
	}
}

func main() {
	filePath := "./input"
	g := getGameFromFile(filePath)
	score := g.GameScore()
	fmt.Printf("PART1: total score = %v\n", score)

	g.UpdateGameToOutcome()
	score = g.GameScore()
	fmt.Printf("PART2: total score = %v\n", score)
}
