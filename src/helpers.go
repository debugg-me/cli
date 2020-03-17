package main

import (
	"fmt"

	"github.com/fatih/color"
)

func printFailed() {
	c := color.New(color.BgRed).Add(color.Bold).Add(color.FgWhite)
	c.Printf(" FAILED ")
}

func printPassed() {
	c := color.New(color.BgGreen).Add(color.Bold).Add(color.FgWhite)
	c.Printf(" PASSED ")
}

func printStderr(stderr string) {
	c := color.New(color.FgYellow)
	c.Println("[STDERR]:")

	fmt.Println(stderr)
}

func printTitle(title string) {
	c := color.New(color.FgWhite).Add(color.Bold)
	c.Println(title)
}
