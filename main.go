package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Aranaris/dwellstats/internal/cmd"
)

func printPrompt() {
	fmt.Print("dwellstats", "> ")
}

func main() {
	cl, err := cmd.InitializeCommands()
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewScanner(os.Stdin)

	a := "5500 Grand Lake Dr, San Antonio, TX, 78244"
	
	fmt.Printf("Default address: %v\n", a)

	printPrompt()

	exec:
	for reader.Scan() {
		text := reader.Text()
		fmt.Println("")

		err = cl.HandleCommand(text)
		if err != nil {
			fmt.Println(err)
		}
		if text == "exit" {
			break exec
		}
		printPrompt()
	}

}
