package cmd

import (
	"fmt"
	"os/exec"

	//"os/exec"
	"strings"
)

func (e *Environment) ExceShell() error {
	fmt.Println("executing shell command:", e.Shell.Command)

	args := strings.Split(e.Shell.Command, " ")

	cmd := exec.Command(args[0], args[1:]...)
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
