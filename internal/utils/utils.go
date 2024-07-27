package utils

import "fmt"

func PrintError(message string, err error) {
	if err != nil {
		fmt.Printf("%s: %v\n", message, err)
	}
}

func IsValidCommand(command string, validCommands map[string]func(args []string)) bool {
	_, ok := validCommands[command]
	return ok
}
