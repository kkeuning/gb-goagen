# gb-goagen
A gb plugin to run the goagen tool in a gb project.

The goagen generator is used to working with a $GOPATH.

## Installation

    go get github.com/kkeuning/gb-goagen/...

## Pre-reqs

	go get github.com/constabulary/gb/...
	go get github.com/goadesign/goa/goagen
	
This is based on Doug Clark's gb-run plug-in, tailored for running goagen specifically.

Usage:
  gb goagen [arguments]

Either run gb-goagen from inside your desired output directory, or
specify the output directory with -o.

Example running from GB_PROJECT_DIR (see gb env):
	gb goagen bootstrap -d goa-adder/design -o $PWD/src/goa-adder

Runs as if you had executed the following:
	GOPATH=$GOPATH:$PWD goagen bootstrap -d goa-adder/design -o $PWD/src/goa-adder

Getting help:

	Same as gb -h, shows gb help, NOT goagen help:
		gb goagen -h

	To see this help together with goagen help:
		gb goagen

	To see just the goagen help:
		goagen
		goagen -h
