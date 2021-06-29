package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("error missing required parameters (./go-envdir ./name_dir_env_file command [arg...]")
		return
	}

	path := os.Args[1]
	cmd := os.Args[2:]
	evn, err := ReadDir(path)
	if err != nil {
		log.Fatal(err)
		return
	}

	exitCode := RunCmd(cmd, evn)
	os.Exit(exitCode)
}
