package main

import (
	"bytes"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type AdditionalCommand struct {
	Command string   `json:"command"`
	Output  []string `json:"output"`
}

func checkAdditionalCommands() []AdditionalCommand {
	commands := []AdditionalCommand{}

	// here we can check for commands that are set as ENV variables.
	// The env variable should be named SIMPLESTATSCOMMANDX where X is 0 - 9
	//
	// if those env variables are set, this function will run them and populate
	// the output in the JSON payload
	for index := 0; index < 10; index++ {
		envvar := os.Getenv("SIMPLESTATSCOMMAND" + strconv.Itoa(index))
		if envvar == "" {
			continue
		}

		commandParts := strings.Split(envvar, " ")
		var result bytes.Buffer

		// get the output of the command
		commandRunner := exec.Command(commandParts[0])
		if len(commandParts) > 1 {
			commandRunner = exec.Command(commandParts[0], commandParts[1:]...)
		}
		commandRunner.Stdout = &result
		commandRunner.Run()

		output := strings.Split(result.String(), "\n")

		currentCommand := AdditionalCommand{
			Command: envvar,
			Output:  output[0 : len(output)-1],
		}
		commands = append(commands, currentCommand)
	}

	return commands
}
