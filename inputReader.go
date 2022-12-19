package main

import (
	term "github.com/nsf/termbox-go"
)

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
			//fallthrough
			return
		case term.KeyArrowRight:
			s = "RIGHT"
			return
		default:
			term.Sync()
		}
	}
}
