package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func timer(duration *time.Duration) func() {
	start := time.Now()
	return func() {
		*duration = time.Since(start)
	}
}

func generate[U uint8 | uint16 | uint32]() (generateDuration time.Duration) {
	// create the .go file which finds if a number is even or odd through if-else-s
	defer timer(&generateDuration)()
	generatedFile, err := os.Create("unholy.go")
	if err != nil {
		fmt.Println("Error while trying to create the go file with error: ", err)
		return
	}

	startString := `package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	number, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Rerun the program with a valid number!")
		return
	}
	if number == 0 {
		fmt.Println("Even")
	} else if number == 1 {
		fmt.Println("Odd")
	}`
	// first if else has to be hardcoded because it doesn't fit the else-if patterns that's looped
	generatedFile.WriteString(startString)

	// add all possible conditions
	elseIfTemplate := ` else if number == %d {
		fmt.Println("Even")
	} else if number == %d {
		fmt.Println("Odd")
	}`
	// overflow back to 0 indicates all numbers covered, incrementing by 2 since writing both even and odd at once
	for i := U(2); i != U(0); i += U(2) {
		generatedFile.WriteString(fmt.Sprintf(elseIfTemplate, i, i+1))
	}

	endString := `
}`
	generatedFile.WriteString(endString)
	return
}

func build() (buildDuration time.Duration) {
	defer timer(&buildDuration)()
	goCommandString := "go"
	buildString := "build"
	fileNameString := "unholy.go"
	cmd := exec.Command(goCommandString, buildString, fileNameString)
	stdoutBytes, err := cmd.Output()
	if err != nil {
		fmt.Println("\tErrored while building the binary with error: ", err)
		return
	}
	if len(stdoutBytes) != 0 {
		fmt.Println("\tThe output while building the binary was: ", string(stdoutBytes))
	}
	return
}
