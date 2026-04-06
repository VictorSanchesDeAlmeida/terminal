package executor

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"
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
		// Replace shell process with the target command so Ctrl+C reaches it directly.
		cmd = exec.Command("sh", "-c", "exec "+input)

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

	sigCh := make(chan os.Signal, 4)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	defer signal.Stop(sigCh)

	stopForward := make(chan struct{})
	interruptCount := 0
	go func() {
		for {
			select {
			case sig := <-sigCh:
				if sig == os.Interrupt {
					interruptCount++
					sendSignalToProcessTree(cmd, syscall.SIGINT)
					if interruptCount > 1 {
						sendSignalToProcessTree(cmd, syscall.SIGKILL)
					}
					continue
				}

				sendSignalToProcessTree(cmd, syscall.SIGTERM)
				time.Sleep(300 * time.Millisecond)
				sendSignalToProcessTree(cmd, syscall.SIGKILL)
			case <-stopForward:
				return
			}
		}
	}()

	err := cmd.Wait()
	close(stopForward)

	return err
}

func sendSignalToProcessTree(cmd *exec.Cmd, sig syscall.Signal) {
	if cmd == nil || cmd.Process == nil {
		return
	}

	pids := []int{cmd.Process.Pid}
	pids = append(pids, listDescendantPIDs(cmd.Process.Pid)...)

	for i := len(pids) - 1; i >= 0; i-- {
		_ = syscall.Kill(pids[i], sig)
	}
}

func listDescendantPIDs(rootPID int) []int {
	if runtime.GOOS != "darwin" && runtime.GOOS != "linux" {
		return nil
	}

	visited := map[int]struct{}{}
	var result []int
	queue := []int{rootPID}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		children := listDirectChildren(current)
		for _, child := range children {
			if _, ok := visited[child]; ok {
				continue
			}

			visited[child] = struct{}{}
			result = append(result, child)
			queue = append(queue, child)
		}
	}

	return result
}

func listDirectChildren(pid int) []int {
	cmd := exec.Command("pgrep", "-P", strconv.Itoa(pid))
	output, err := cmd.Output()
	if err != nil {
		return nil
	}

	lines := strings.Fields(string(output))
	children := make([]int, 0, len(lines))
	for _, line := range lines {
		childPID, convErr := strconv.Atoi(strings.TrimSpace(line))
		if convErr != nil {
			continue
		}
		children = append(children, childPID)
	}

	return children
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
