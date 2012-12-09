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
	var path string
	python := os.Getenv("PYTHON")

	// if $PYTHON is a path to an executable, use it; if not, try to expand the
	// shortcuts listed above
	if _, err := os.Stat(python); err == nil {
		path = python
	} else {
		path, _ = versions[python]
	}

	// make sure we have something to exec
	if _, err := os.Stat(path); err != nil {
		log.Fatalf("unknown python version '%s'\n", python)
	}

	// give python our arguments but the path to the python executable
	os.Args[0] = path

	// start python connected to our fds 0, 1, and 2
	p, err := os.StartProcess(path, os.Args, &os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	})
	if err != nil {
		log.Fatalf("failed to start python: %v\n", err)
	}

	// wait for it to execute, then clean up
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
