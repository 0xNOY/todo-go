package router

import (
	"flag"
	"fmt"
	"os"
)

type Command struct {
	Name        string
	Function    func(*flag.FlagSet)
	Description string
	FlagSet     *flag.FlagSet
}

type Router struct {
	Commands []Command
}

func (r *Router) Start() {
	if len(os.Args) > 1 {
		for _, command := range r.Commands {
			if command.Name == os.Args[1] {
				command.Function(command.FlagSet)
				return
			}
		}
	}

	for _, command := range r.Commands {
		fmt.Println(command.Name, command.Description)
	}
}

func (r *Router) AddCommand(name string, description string, function func(*flag.FlagSet)) {
	r.Commands = append(r.Commands, Command{
		Name:        name,
		Function:    function,
		Description: description,
		FlagSet:     flag.NewFlagSet(name, flag.ExitOnError),
	})
}
