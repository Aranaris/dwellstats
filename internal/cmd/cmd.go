package cmd

import (
	"errors"
	"fmt"
	"sync"
)

type Command struct{
	Name string
	Description string
	Config *APIConfig
}

type CommandList map[string]Command

type APIConfig struct {
	TargetURL string
	Mutex *sync.RWMutex
	PropData map[*string]struct{}
}

func InitializeCommands() (*CommandList, error) {

	cl := make(CommandList)

	cfg := APIConfig{
		TargetURL: "127.0.0.1:1337",
	}

	Exit := Command{
		Name: "exit",
		Description: "Exits out of the program",
	}

	Help := Command{
		Name: "help",
		Description: "prints out all supported commands",
		Config: &cfg,
	}

	cl[Help.Name] = Help
	cl[Exit.Name] = Exit

	return &cl, nil
}


func (cl *CommandList) CommandHelp() error {
	fmt.Println("List of all valid commands:")
	for _, v := range *cl {
		fmt.Printf("Command: %s || Description: %s", v.Name, v.Description)
		fmt.Println("")
	}
	return nil
}

func (cl *CommandList) CommandExit() error {
	fmt.Println("Closing...")
	return nil
}

func (cl *CommandList) HandleCommand(input string) error {
	if input == "help" {
		cl.CommandHelp()
		return nil
	}

	if input == "exit" {
		cl.CommandExit()
		return nil
	}
	
	return errors.New("Command not found: " + input)
}
