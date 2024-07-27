package utils

import "fmt"

const (
	Red   = "\033[31m"
	Reset = "\033[0m"
)

func PrintError(message string, err error) {
	if err != nil {
		fmt.Printf("%s%s: %v%s\n", Red, message, err, Reset)
	}
}
