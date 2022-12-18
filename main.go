package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	term "github.com/nsf/termbox-go"
)

// LoadMaze read maze into string array
func LoadMaze(path string) ([]string, error) {
	mz := []string{}
	fileHandler, e := os.Open(path)
	if e != nil {
		return nil, e
	}
	defer fileHandler.Close()

	scanner := bufio.NewScanner(fileHandler)
	for scanner.Scan() {
		ln := scanner.Text()
		mz = append(mz, ln)
	}
	return mz, nil
}

// ReadInputSpecialCharacters read input special characters: esc, up, down, left, right
func ReadInputSpecialCharacters() (s string, e error) {
	e = term.Init()
	defer term.Close()
	if e != nil {
		return
	}
	ev := term.PollEvent()

	switch ev.Key {
	case term.KeyEsc:
		s = "ESC"
	case term.KeyArrowUp:
		s = "UP"
	case term.KeyArrowDown:
		s = "DOWN"
	case term.KeyArrowLeft:
		s = "LEFT"
	case term.KeyArrowRight:
		s = "RIGHT"
	}

	return
}

// PrintMaze printing maze to cmd
func PrintMaze(maze []string) {
	for _, line := range maze {
		fmt.Println(line)
	}
}

func main() {
	mazeMap, e := LoadMaze("maze01.txt")
	if e != nil {
		log.Println("Failed to load maze:", e)
		return
	}

	for {
		PrintMaze(mazeMap)
		input, e := ReadInputSpecialCharacters()
		fmt.Println(input, e)
		break
	}
}
