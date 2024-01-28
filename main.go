package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cleanScreen() {
	fmt.Print("\033[H\033[2J")
}

func readIntFromConsole(reader *bufio.Reader) (int, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	// Trim the newline character from the input
	input = strings.TrimSpace(input)

	// Parse the string into an integer
	number, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return number, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	cleanScreen()

	printMenu()
	initLaunchers()

	for {
		fmt.Println("Enter your choice: ")
		choice, _ := readIntFromConsole(reader)

		if item, ok := menu[menuOption(choice)]; ok {
			item.action()
		} else {
			cleanScreen()
			fmt.Println("Invalid choice")
			printMenu()
		}
	}
}
