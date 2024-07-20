package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func cmdExit(args []string) {
	println("Exiting shell")
	os.Exit(0)
}
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

func cmdTouch(args []string) {
	if len(args) == 0 {
		fmt.Println("touch: missing file operand")
		return
	}
	for _, filename := range args {
		_, err := os.Stat(filename)
		if os.IsNotExist(err) {
			// file doesnot exist create it
			file, err := os.Create(filename)
			if err != nil {
				fmt.Printf("touch: cannot create '%s': %v\n", filename, err)
				continue
			}
			file.Close()
			fmt.Printf("Created file: %s\n", filename)
		} else if err == nil {
			// file exists update its timestamp
			now := time.Now()
			err := os.Chtimes(filename, now, now)
			if err != nil {
				fmt.Printf("touch: cannot touch '%s': %v\n", filename, err)
				continue
			} else {
				fmt.Printf("touch: error checking '%s': %v\n", filename, err)
			}

		}

	}
}
func cmdPwd(args []string) {}

var validCommands = map[string]func(args []string){
	"echo":  cmdEcho,
	"pwd":   cmdPwd,
	"about": cmdAbout,
	"exit":  cmdExit,
	"touch": cmdTouch,
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
