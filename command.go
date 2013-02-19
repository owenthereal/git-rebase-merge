package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

type Command struct {
	input string
}

func (c *Command) Run() {
	name, args := c.parse(c.input)

	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func (c *Command) parse(input string) (string, []string) {
	inputs := strings.Split(input, " ")
	name := inputs[0]
	args := inputs[1:]

	return name, args
}
