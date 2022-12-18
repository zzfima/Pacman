package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

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

	for {
		switch ev := term.PollEvent(); ev.Key {
		case term.KeyEsc:
			s = "ESC"
			return
		case term.KeyArrowUp:
			s = "UP"
			return
		case term.KeyArrowDown:
			s = "DOWN"
			return
		case term.KeyArrowLeft:
			s = "LEFT"
			return
		case term.KeyArrowRight:
			s = "RIGHT"
			return
		default:
			term.Sync()
		}
	}
}

// PrintMaze printing maze to cmd
func PrintMaze(maze []string) {
	for _, line := range maze {
		fmt.Println(line)
	}
}

// CleanScreen clean all info from screen
func CleanScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Run()
}

func main() {
	mazeMap, e := LoadMaze("maze01.txt")
	if e != nil {
		log.Println("Failed to load maze:", e)
		return
	}

	CleanScreen()
	PrintMaze(mazeMap)
	for {
		input, e := ReadInputSpecialCharacters()
		fmt.Println(input)

		if input == "ESC" || e != nil {
			break
		}

		if input == "UP" || input == "DOWN" || input == "LEFT" || input == "RIGHT" {
			PrintMaze(mazeMap)
		}

		term.Sync()
	}
}
