package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cmdAbout(args []string) {
	fmt.Println("goshell")
	fmt.Println("A simple POSIX-compliant shell implementation in Go.")
	fmt.Println("GitHub repository: [GoShell](https://github.com/mnsdojo/goshell)")
}
func cmdEcho(args []string) {
	if len(args) == 0 {
		println()
	} else {
		fmt.Println(strings.Join(args, ""))
	}
}

func cmdPwd(args []string) {}

var validCommands = map[string]func(args []string){

	"echo":  cmdEcho,
	"pwd":   cmdPwd,
	"about": cmdAbout,
}

func isValidCommand(command string) bool {
	_, ok := validCommands[command]
	return ok
}

// runshell starts the shell

func RunShell() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$ ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		parts := strings.Fields(input)
		command := parts[0]
		args := parts[1:]

		cmdFn, ok := validCommands[command]
		if !ok {
			fmt.Printf("%s: command not found\n", command)
			continue
		}
		cmdFn(args)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading standard input: %v\n", err)
	}
}
