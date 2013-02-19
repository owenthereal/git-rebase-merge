package main

var commands = []Command{
	Command{"git help"},
	Command{"ls -all"},
}

var commander = Commander{commands}

func main() {
	commander.RunAll()
}
