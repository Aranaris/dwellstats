package cmd

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/Aranaris/dwellstats/internal/models"
	"github.com/Aranaris/dwellstats/internal/request"
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
	PropData map[*string]models.Property
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
	}

	Find := Command{
		Name: "find",
		Description: "given an address, retrieve property data",
		Config: &cfg,
	}

	cl[Help.Name] = Help
	cl[Exit.Name] = Exit
	cl[Find.Name] = Find

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

func (cl *CommandList) CommandFind(address string) error {
	cfg := (*cl)["find"].Config

	fmt.Printf("Searching for %s...", address)

	p, err := request.GetPropertyByAddress(cfg.TargetURL, address)
	if err != nil {
		return err
	}

	cfg.PropData[&address] = *p
	fmt.Println(p)

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

	s := strings.SplitN(input, " ", 2)
	if s[0] == "find" {
		if len(s) < 2 {
			return errors.New("no address provided")
		}
		err := cl.CommandFind(s[1])
		if err != nil {
			return err
		}
	}
	
	return errors.New("Command not found: " + input)
}
