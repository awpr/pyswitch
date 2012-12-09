Arch Linux is admirably brave to install Python 3 and only Python 3 by default,
but that choice has a tendency to break a lot of software that expects "python"
to be Python 2.  Many Arch users are no doubt familiar with the pain of
replacing all occurrences of "python" with "python2" in build scripts.

This software's purpose is to circumvent this.  It chooses the python
executable to act as based on the environment variable PYTHON:

- if empty or unset, "/usr/bin/python"
- if "2", "/usr/bin/python2"
- if "3", "/usr/bin/python3"
- if anything else, fail and print an error to stderr

It was made while trying to compile llvm from source on Arch, and it fulfills
that purpose successfully: "PYTHON=2 make"

It should not interfere with any use of python when the environment variable is
not set.

To use:

- go get github.com/awpr/pyswitch
- ln -s $GOROOT/bin/pyswitch $GOROOT/bin/python
- ensure that $GOROOT/bin is ahead of /usr/bin in $PATH
