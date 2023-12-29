package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func TestInvalidProxyAddress(t *testing.T) {
	executer := DefaultExecuter{}
	err := executer.ExecuteCommand("invalid-proxy-address", "echo", []string{"Hello World"})
	if err == nil || !strings.Contains(err.Error(), "failed to connect to proxy") {
		t.Errorf("Expected proxy connection failure, got %v", err)
	}
}

func TestExecuteCommand_InvalidProxyAddress(t *testing.T) {
	executer := DefaultExecuter{}
	err := executer.ExecuteCommand("invalid-proxy", "echo", []string{"Hello World"})
	if err == nil || !strings.Contains(err.Error(), "failed to connect to proxy") {
		t.Errorf("Expected proxy connection failure, got %v", err)
	}
}

func TestTerminateProcess(t *testing.T) {
	cmd := exec.Command("sleep", "10")

	if err := cmd.Start(); err != nil {
		t.Fatalf("Failed to start process: %v", err)
	}

	if cmd.Process == nil {
		t.Fatalf("Process was not started")
	}

	terminateProcess(cmd)

	time.Sleep(time.Second)

	if err := cmd.Process.Signal(nil); err == nil {
		t.Errorf("Process was not terminated")
	}
}

func TestClearLine(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}

	oldStdout := os.Stdout
	os.Stdout = w
	defer func() { os.Stdout = oldStdout }()

	clearLine()

	w.Close()

	var buf bytes.Buffer
	buf.ReadFrom(r)

	expected := "\r" + strings.Repeat(" ", 50) + "\n"
	if got := buf.String(); got != expected {
		t.Errorf("clearLine() = %q, want %q", got, expected)
	}
}
