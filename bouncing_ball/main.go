package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	const (
		width  = 50
		height = 10

		cellEmpty = ' '
		cellBall  = '⚽'

		maxFrames = 1200
		speed     = time.Second / 30
	)

	var (
		px, py int
		vx, vy = 1, 1
		cell   rune
	)
	screen.Clear()

	buf := make([]rune, 0, width*height)
	board := make([][]bool, width)
	for i := range board {
		board[i] = make([]bool, height)
	}

	for i := 0; i < maxFrames; i++ {
		px += vx
		py += vy

		if px <= 0 || px >= width-1 {
			vx *= -1
		}

		if py <= 0 || py >= height-1 {
			vy *= -1
		}

		for y := range board[0] {
			for x := range board {
				board[x][y] = false
			}
		}

		buf = buf[:0]
		board[px][py] = true
		// draw the board
		for y := range board[0] {
			for x := range board {
				cell = cellEmpty
				if board[x][y] {
					cell = cellBall
				}
				buf = append(buf, cell, ' ')
			}
			buf = append(buf, '\n')
		}
		time.Sleep(speed)
		screen.MoveTopLeft()
		fmt.Printf(string(buf))
	}
}
