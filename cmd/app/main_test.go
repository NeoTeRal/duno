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
	out, err := exec.Command("pwd").Output()
	if err != nil {
		return 
	}
	
	fmt.Printf("%s\n", out)
}
