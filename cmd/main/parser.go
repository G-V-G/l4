package main

import (
	"regexp"
	"strings"
	"github.com/G-V-G/l4/commands"
	"github.com/G-V-G/l4/engine"
)

func parse(str string) (cmd engine.Command) {
	if len(str) == 0 {
		errStr := "Syntax error: empty string"
		cmd = &commands.PrintCommand{Arg: errStr}
		return
	}
	includesCommand := `^(print|reverse) `
	isValidCommand, _ := regexp.MatchString(includesCommand, str)
	if !isValidCommand {
		errStr := "Syntax error: unknown command"
		cmd = &commands.PrintCommand{Arg: errStr}
		return
	}
	validTemplate := includesCommand + `\S+$`
	isValidArgs, _ := regexp.MatchString(validTemplate, str)
	if !isValidArgs {
		errStr := "Syntax error: invalid argument"
		cmd = &commands.PrintCommand{Arg: errStr}
		return
	}
	parts := strings.Fields(str)
	if parts[0] == "print" {
		cmd = &commands.PrintCommand{Arg: parts[1]}
	} else {
		cmd = &commands.ReverseCommand{Arg: parts[1]}
	}
	return
}
