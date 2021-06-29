package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	if len(cmd) == 0 {
		fmt.Println("no cmd")
		return 1
	}
	c := exec.Command(cmd[0], cmd[1:]...) // nolint
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	for k, p := range env {
		if p.NeedRemove {
			os.Unsetenv(k)
		}
		err := os.Setenv(k, p.Value)
		if err != nil {
			fmt.Printf("error setting environment variable %s = %v\n", k, p.Value)
		}
	}

	if err := c.Run(); err != nil {
		log.Fatal(err)
	}

	return c.ProcessState.ExitCode()
}
