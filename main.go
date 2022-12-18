package main

import (
	"fmt"
	"log"
	"os/exec"

	term "github.com/nsf/termbox-go"
)

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

// Sprite place on maze
type Sprite struct {
	row    int
	column int
}

func main() {
	mazeMap, e := LoadMaze("maze01.txt")
	if e != nil {
		log.Println("Failed to load maze:", e)
		return
	}

	player, e := LoadPlayer(mazeMap)
	fmt.Println(player)
	if e != nil {
		log.Println("Failed to load player:", e)
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
