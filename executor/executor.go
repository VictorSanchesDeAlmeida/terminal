package executor

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func Execute(input string) {
		if input == "" {
		return
	}

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		// cmd.exe é mais garantido que PowerShell
		cmd = exec.Command("cmd", "/C", input)

	case "darwin", "linux":
		// sh é o mais compatível possível
		cmd = exec.Command("sh", "-c", input)

	default:
		fmt.Println("Sistema operacional não suportado")
		return
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Println("Erro:", err)
	}
}
