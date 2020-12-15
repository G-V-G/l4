package commands

import (
	"fmt"
	"github.com/G-V-G/l4/engine"
)

//PrintCommand struct for string output
type PrintCommand struct {
	Arg string
}

//Execute method for print command
func (p *PrintCommand) Execute(_ engine.Handler) {
	fmt.Println(p.Arg)
}