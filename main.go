package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/huavcjj/interpreter-playground/repl"
)

func main() {
	current, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello, %s! Welcome to the interpreter playground.\n", current.Username)
	fmt.Printf("Feel free to type your code below:\n")
	repl.Start(os.Stdin, os.Stdout)
}
