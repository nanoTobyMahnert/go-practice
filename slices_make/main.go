package main

import (
	"strings"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	s.PrintBacking = true
	s.MaxPerLine = 10

	tasks := []string{"jump", "run", "read"}

	upTasks := make([]string, 0, len(tasks))
	s.Show("upTasks", upTasks)

	for _, task := range tasks {
		upTasks = append(upTasks, strings.ToUpper(task))
		s.Show("upTasks", upTasks)
	}

	upTasks = append(upTasks, "PLAY")
	s.Show("upTasks", upTasks)
}
