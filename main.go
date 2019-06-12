package main

import (
	"fmt"
	"os"
	"os/user"
	"zLang/repl"
)

func main() {
	usr, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! this is the Zlang programming language\n", usr.Username)
	fmt.Printf("please type some commands\n")

	repl.Start(os.Stdin, os.Stdout)
}
