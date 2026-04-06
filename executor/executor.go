package executor

import (
	"fmt"
	"os"
	"os/exec"
)

func Execute(input string) {
	if input == "" {
		return
	}

	cmd := exec.Command("powershell", "-Command", input)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Println("Erro:", err)
	}
}
