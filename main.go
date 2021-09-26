package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/MarcBernstein0/the-monkey-programming-lang-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %v. Welcome to the Monkey programming language repl.\n",
		user.Username)

	repl.Start(os.Stdin, os.Stdout)

	fmt.Printf("Goodbye %v!\n", user.Username)

}
