package executor

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
)

func Execute(input string) {
	input = strings.TrimSpace(input)
	if input == "" {
		return
	}

	if handleBuiltin(input) {
		return
	}

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/C", input)

	case "darwin", "linux":
		cmd = exec.Command("sh", "-c", input)

	default:
		fmt.Println("Sistema operacional não suportado")
		return
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := runWithInterruptSupport(cmd); err != nil {
		fmt.Println("Erro:", err)
	}
}

func runWithInterruptSupport(cmd *exec.Cmd) error {
	if err := cmd.Start(); err != nil {
		return err
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	defer signal.Stop(sigCh)

	stopForward := make(chan struct{})
	go func() {
		for {
			select {
			case <-sigCh:
				if cmd.Process == nil {
					continue
				}

				if err := cmd.Process.Signal(os.Interrupt); err != nil {
					_ = cmd.Process.Kill()
				}
			case <-stopForward:
				return
			}
		}
	}()

	err := cmd.Wait()
	close(stopForward)

	return err
}

func PromptPrefix() string {
	currentDir, err := os.Getwd()
	if err != nil {
		return "devshell >> "
	}

	return currentDir + " >> "
}

func handleBuiltin(input string) bool {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return true
	}

	switch parts[0] {
	case "exit", "quit":
		os.Exit(0)
		return true
	case "pwd":
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Erro:", err)
			return true
		}
		fmt.Println(cwd)
		return true
	case "cd":
		target, err := resolveCdTarget(parts)
		if err != nil {
			fmt.Println("Erro:", err)
			return true
		}

		if err := os.Chdir(target); err != nil {
			fmt.Println("Erro:", err)
		}
		return true
	case "clear":
		fmt.Print("\033[H\033[2J")
		return true
	}

	return false
}

func resolveCdTarget(parts []string) (string, error) {
	if len(parts) == 1 || parts[1] == "~" {
		return os.UserHomeDir()
	}

	target := parts[1]

	if strings.HasPrefix(target, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		target = filepath.Join(home, strings.TrimPrefix(target, "~/"))
	}

	return target, nil
}
