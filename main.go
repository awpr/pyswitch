package main

import (
	"log"
	"os"
)

var versions = map[string]string {
	"2": "/usr/bin/python2",
	"3": "/usr/bin/python3",
	"":  "/usr/bin/python",
}

func main() {
	python := os.Getenv("PYTHON")
	path, _ := versions[python]

	if _, err := os.Stat(path); err != nil {
		log.Fatalf("unknown python version '%s'", python)
	}

	args := os.Args
	args[0] = path

	p, err := os.StartProcess(path, args, &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	})
	if err != nil {
		log.Fatalf("failed to start python: %v\n", err)
	}

	s, err := p.Wait()
	if err != nil {
		log.Fatalf("failed to wait for python: %v\n", err)
	}

	if !s.Success() {
		log.Fatalf("python did not exit successfully: %v\n", s.Sys())
	}
}
