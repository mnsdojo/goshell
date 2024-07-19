package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$ ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		fmt.Println("You entered:", input)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "err")
	}

}
