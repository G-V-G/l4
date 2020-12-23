package main

import (
	"bufio"
	"os"

	"github.com/G-V-G/l4/engine"
)

func main() {
	filename := "commands.txt"
	eventLoop := new(engine.Engine)
	eventLoop.Start()
	if input, err := os.Open(filename); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := parse(commandLine)
			eventLoop.Post(cmd)
		}
	} else {
		panic(err)
	}
	eventLoop.AwaitFinish()
}
