package main

import (
	"devshell/completer"
	"devshell/executor"
	"os"

	"github.com/c-bata/go-prompt"
)

func main() {

	currentDir, _ := os.Getwd()
	p := prompt.New(
		executor.Execute,
		completer.Complete,
		prompt.OptionPrefix(currentDir+" >> "),
	)

	p.Run()
}
