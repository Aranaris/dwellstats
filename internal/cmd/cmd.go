package cmd

import "sync"

type Command struct{
	Name string
	Description string
	Config *APIConfig
}

type CommandList map[string]Command

type APIConfig struct {
	NextURL string
	PreviousURL string
	Mutex *sync.RWMutex
	PropData map[*string]struct{}
}
