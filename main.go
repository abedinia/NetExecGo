package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/fatih/color"
)

type CommandExecuter interface {
	ExecuteCommand(proxyAddr, command string, args []string) error
}

type DefaultExecuter struct{}

func (e DefaultExecuter) ExecuteCommand(proxyAddr, command string, args []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Minute)
	defer cancel()

	dialer, err := net.Dial("tcp", proxyAddr)
	if err != nil {
		return fmt.Errorf("failed to connect to proxy: %v", err)
	}
	defer dialer.Close()

	fmt.Printf("Using proxy IP: %s\n", proxyAddr)
	fmt.Println("..............................")

	cmd := exec.CommandContext(ctx, command, args...)
	cmd.Env = append(os.Environ(), "ALL_PROXY=socks5://"+proxyAddr)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("error creating stdout pipe: %v", err)
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("error creating stderr pipe: %v", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("error starting command: %v", err)
	}

	handleSignals(cmd, cancel)

	displayPandaEmojis(ctx)

	manageCommandOutput(stdoutPipe, false)
	manageCommandOutput(stderrPipe, true)

	return awaitCommandCompletion(cmd, ctx)
}

func displayPandaEmojis(ctx context.Context) {
	pandaEmojis := []string{"⏳", "⌛"}
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				for _, panda := range pandaEmojis {
					fmt.Printf("\r%s ", panda)
					time.Sleep(200 * time.Millisecond)
				}
			}
		}
	}()
}

func handleSignals(cmd *exec.Cmd, cancel context.CancelFunc) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signals
		fmt.Println("\nInterrupt received, stopping...")
		cancel()
		terminateProcess(cmd)
	}()
}

func manageCommandOutput(pipe io.Reader, isStderr bool) {
	go func() {
		scanner := bufio.NewScanner(pipe)
		for scanner.Scan() {
			line := scanner.Text()
			clearLine()
			if isStderr {
				color.New(color.FgHiBlack).Println(line)
			} else {
				color.Green(line)
			}
		}
	}()
}

func awaitCommandCompletion(cmd *exec.Cmd, ctx context.Context) error {
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case <-ctx.Done():
		terminateProcess(cmd)
		return ctx.Err()
	case err := <-done:
		clearLine()
		if err != nil {
			return fmt.Errorf("command finished with error: %v", err)
		}
	}

	clearLine()
	return nil
}

func terminateProcess(cmd *exec.Cmd) {
	if cmd.Process != nil {
		cmd.Process.Kill()
	}
}

func clearLine() {
	fmt.Printf("\r%s\n", strings.Repeat(" ", 50))
}

func handleFatalError(message string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}


func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go [proxy] [command] [command arguments]")
		os.Exit(1)
	}

	executer := DefaultExecuter{}
	err := executer.ExecuteCommand(os.Args[1], os.Args[2], os.Args[3:])
	if err != nil {
		fmt.Println(err)
	}
}
