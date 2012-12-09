package main

import (
	"log"
	"os"
	"syscall"
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
		log.Fatalf("unknown python version '%s'\n", python)
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
		// try to exit with the right exit code
		if ws, ok := s.Sys().(syscall.WaitStatus); ok {
			os.Exit(ws.ExitStatus())
		}

		// not on a Posix system, so we don't know the exit code.  1's as
		// good as any other
		os.Exit(1)
	}
}
