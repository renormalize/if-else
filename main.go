package main

import (
	"fmt"
	"os"
	"time"
)

// Inspired by https://andreasjhkarlsson.github.io//jekyll/update/2023/12/27/4-billion-if-statements.html
func main() {
	if len(os.Args) != 3 {
		printUsage()
		return
	}
	subCommand := os.Args[1]
	bitWidth := os.Args[2]
	if subCommand != "generate" && subCommand != "build" {
		printUsage()
		return
	}
	if bitWidth != "8" && bitWidth != "16" && bitWidth != "32" {
		printUsage()
		return
	}

	fmt.Println()
	fmt.Println("\tGenerating the super efficient program to find if a number is even or odd!")
	var generateDuration time.Duration
	switch bitWidth {
	case "8":
		generateDuration = generate[uint8]()
	case "16":
		generateDuration = generate[uint16]()
	case "32":
		generateDuration = generate[uint32]()
	default:
		printUsage()
		return
	}
	fmt.Println("\tFile generated in ", generateDuration, "!")

	if subCommand == "build" {
		var buildDuration time.Duration
		fmt.Println("\tBuilding the super efficient program to find if a number is even or odd!")
		buildDuration = build()
		fmt.Println("\tFile built in ", buildDuration, "!")
	}
	fmt.Println()
}

func printUsage() {
	fmt.Println()
	fmt.Println("\tWrong usage, first argument must be \"generate\" or \"build\" and second argument must be bit width")
	fmt.Println("\t\"generate\" generates the .go file")
	fmt.Println("\t\"build\" generates the .go file and builds it")
	fmt.Println("\tBit width must be 8, 16 or 32")
	fmt.Println()
}
