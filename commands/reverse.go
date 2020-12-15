package commands

import "github.com/G-V-G/l4/engine"

//ReverseCommand struct for string reversion
type ReverseCommand struct {
	Arg string
}

//Execute method for reverse command
func (r *ReverseCommand) Execute(loop engine.Handler) {
	reversedStr := ""
	for _, char := range r.Arg {
		reversedStr = string(char) + reversedStr
	}
	loop.Post(&PrintCommand{Arg: reversedStr})
}