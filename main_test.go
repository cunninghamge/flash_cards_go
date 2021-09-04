package main

import (
	"errors"
	"os"
	"os/exec"
	"testing"
)

func TestExitWithError(t *testing.T) {
	if os.Getenv("OS_EXIT_CALLED") == "1" {
		exitWithError(errors.New("error"))
		return
	}
	subTest := exec.Command(os.Args[0], "-test.run=TestNewRoundExit")
	subTest.Env = append(os.Environ(), "OS_EXIT_CALLED=1")
	err := subTest.Run()
	if exitError, ok := err.(*exec.ExitError); !ok || exitError.Success() {
		t.Error("process exited with no error, wanted exit status 1")
	}
}
