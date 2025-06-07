package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"internal/cmd"
)

func printPrompt() {
	fmt.Print("dwellstats", "> ")
}

func main() {
	_, err := cmd.InitializeCommands()
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewScanner(os.Stdin)
	printPrompt()

	exec:
	for reader.Scan() {
		text := reader.Text()
		fmt.Println("")

		if text == "exit" {
			break exec
		}
		printPrompt()
	}

}
