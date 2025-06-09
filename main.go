package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"internal/cmd"

	"internal/models"
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

	// a := "5500 Grand Lake Dr, San Antonio, TX, 78244"
	
	sample := models.Address{}
	sample.Street = "5500 Grand Lake Dr"
	sample.City = "San Antonio"
	sample.State = "TX"
	sample.Zip = "78244"
	fmt.Printf("Default address: %v\n", sample)

	printPrompt()

	exec:
	for reader.Scan() {
		text := reader.Text()
		fmt.Println("")

		cl.HandleCommand(text)
		if text == "exit" {
			break exec
		}
		printPrompt()
	}

}
