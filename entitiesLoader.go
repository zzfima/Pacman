package main

import (
	"bufio"
	"fmt"
	"os"
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

// LoadPlayer load player from maze
func LoadPlayer(maze []string) (Sprite, error) {
	for rowIndex, row := range maze {
		for columnIndex, element := range row {
			if element == 'P' {
				return Sprite{rowIndex, columnIndex}, nil
			}
		}
	}

	return Sprite{}, fmt.Errorf("no player found")
}
