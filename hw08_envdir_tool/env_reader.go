package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"strings"
)

type Environment map[string]EnvValue

type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
func ReadDir(dir string) (Environment, error) {
	environment := make(Environment)
	d, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range d {
		name := strings.ReplaceAll(file.Name(), "=", "")

		if file.Size() == 0 {
			environment[name] = EnvValue{"", true}
			continue
		}

		file, err := os.Open(dir + "/" + name)
		if err != nil {
			return nil, err
		}

		s := bufio.NewScanner(file)
		s.Scan()
		val := string(bytes.ReplaceAll([]byte(s.Text()), []byte("\x00"), []byte("\n")))
		val = strings.TrimRight(val, " ")
		val = strings.TrimRight(val, "\t")

		environment[name] = EnvValue{val, false}

		if err = file.Close(); err != nil {
			return nil, err
		}
	}

	return environment, nil
}
