package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	prev := time.Now()
	for scanner.Scan() {
		now := time.Now()
		line := scanner.Text()

		printLine(line, prev, now)

		prev = now
	}
}

var formatSize = fmt.Sprint(len(time.RFC3339Nano))

func printLine(line string, prev, now time.Time) {
	duration := now.Sub(prev)
	fmt.Printf(
		"%-15s %-"+formatSize+"s | %s\n",
		duration,
		now.Format(time.RFC3339Nano),
		line,
	)
}
