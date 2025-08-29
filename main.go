package main

import (
	"flag"
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/leoviggiano/daily/cmd"
	"github.com/leoviggiano/daily/pkg/daily"
)

var (
	currentDaily *daily.Daily
)

func main() {
	flag.Parse()

	currentDaily = daily.NewDaily()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("No arguments provided")
		return
	}

	commands := cmd.Commands(currentDaily)
	cmdMap := make(map[string]cmd.Command)
	for _, command := range commands {
		for _, name := range command.Name() {
			cmdMap[name] = command
		}
	}

	currentCommand, ok := cmdMap[args[0]]
	if !ok {
		fmt.Println("Command not found")
		return
	}

	for _, arg := range args {
		if slices.Contains(cmd.HelpFlags, arg) {
			fmt.Println(currentCommand.Help())
			return
		}
	}

	args = args[1:]
	if len(args) == 1 {
		args = strings.Split(args[0], ",")
	}

	err := currentCommand.Exec(args...)
	if err != nil {
		log.Fatal(err)
	}
}
