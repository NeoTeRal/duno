package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestMain(m *testing.M) {
	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestMainApp(t *testing.T) {
	fmt.Printf("%s\n", string(exec.Command("pwd").Output()))	
}
