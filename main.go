package main

import (
	"bufio"
	"os"
)

var mazeMap []string

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

func main() {
	mazeMap, e := LoadMaze("maze01.txt")
	if e != nil {
		return
	}
	mazeMap[0] = ""
}