package shell

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/mnsdojo/goshell/internal/utils"
)

func cmdLs(args []string) {
	var dir string
	if len(args) == 0 {
		dir = "."
	} else {
		dir = args[0]
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("ls: cannot access '%s' : %v\n", dir, err)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Printf("%s/\n", entry.Name())
		} else {
			fmt.Println(entry.Name())
		}
	}
}

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

func cmdMkdir(args []string) {
	if len(args) == 0 {
		fmt.Println("mkdir: missing file operand")
		return
	}
	for _, dir := range args {
		_, err := os.Stat(dir)
		if err == nil {
			fmt.Printf("mkdir : cannot create directory %s: File exists\n", dir)
			continue
		}
		if os.IsNotExist(err) {
			err := os.Mkdir(dir, 0755)
			// o755 - Permission (rwx)
			if err != nil {
				fmt.Printf("mkdir: cannot create directory '%s' : %v\n", dir, err)
			} else {
				fmt.Printf("Created directory successfully: %s\n", dir)
			}
		} else {
			fmt.Printf("mkdir: cannot access '%s': %v\n", dir, err)
		}

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
				utils.PrintError("touch: cannot create ", err)
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
func cmdPwd(args []string) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("pwd: %v\n", err)
		return
	}
	fmt.Println(dir)
}

func cmdClear(args []string) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Printf("clear: %v\n", err)
	}
}

var validCommands = map[string]func(args []string){
	"echo":  cmdEcho,
	"pwd":   cmdPwd,
	"about": cmdAbout,
	"exit":  cmdExit,
	"touch": cmdTouch,
	"mkdir": cmdMkdir,
	"ls":    cmdLs,
	"clear": cmdClear,
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
		if !isValidCommand(command) {
			fmt.Printf("%s: command not found\n", command)
			continue
		}

		cmdFn := validCommands[command]
		cmdFn(args)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading standard input: %v\n", err)
	}
}
