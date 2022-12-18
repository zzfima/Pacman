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

// ReadInput read input includes esc, arrows etc
func ReadInput() (string, error) {
	e := term.Init()
	defer term.Close()
	if e != nil {
		return "", e
	}
	ev := term.PollEvent()

	if ev.Ch != 0 {
		return string(ev.Ch), nil
	}

	switch ev.Key {
	case term.KeyEsc:
		return "ESC", nil
	case term.KeyArrowUp:
		return "UP", nil
	case term.KeyArrowDown:
		return "DOWN", nil
	case term.KeyArrowLeft:
		return "LEFT", nil
	case term.KeyArrowRight:
		return "RIGHT", nil
	}

	return "", nil
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
		input, e := ReadInput()
		fmt.Println(input, e)
		break
	}
}
