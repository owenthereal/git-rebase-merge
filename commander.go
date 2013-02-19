package main

type Commander struct {
	commands []Command
}

func (c *Commander) RunAll() {
	for _, cmd := range c.commands {
		cmd.Run()
	}
}
