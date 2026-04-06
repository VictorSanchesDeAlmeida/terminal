package main

import (
	"devshell/completer"
	"devshell/executor"

	"github.com/c-bata/go-prompt"
)

func main() {
	p := prompt.New(
		executor.Execute,
		completer.Complete,
		prompt.OptionSwitchKeyBindMode(prompt.EmacsKeyBind),
		prompt.OptionLivePrefix(func() (string, bool) {
			return executor.PromptPrefix(), true
		}),
	)

	p.Run()
}
