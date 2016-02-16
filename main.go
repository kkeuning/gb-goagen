package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var (
	projectroot = os.Getenv("GB_PROJECT_DIR")
	gopath      = os.Getenv("GOPATH")
	args        = os.Args[0:]
)

func main() {
	fatalf := func(format string, args ...interface{}) {
		fmt.Fprintf(os.Stderr, "FATAL: "+format+"\n", args...)
		os.Exit(1)
	}

	switch {
	case len(args) < 2:
		printUsage()
	case projectroot == "":
		fatalf("don't run this binary directly, it is meant to be run as 'gb goagen ...'")
	default:
	}

	env := mergeEnv(os.Environ(), map[string]string{
		"GOPATH": gopath + ":" + projectroot,
	})

	fmt.Println()
	path, err := exec.LookPath("goagen")
	if err != nil {
		fatalf("run: unable to locate %q: %v", args[0], err)
	}

	cmd := exec.Cmd{
		Path: path,
		Args: args,
		Env:  env,

		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	if err := cmd.Run(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	fmt.Println("Done")
}

func mergeEnv(env []string, args map[string]string) []string {
	m := make(map[string]string)
	for _, e := range env {
		v := strings.SplitN(e, "=", 2)
		m[v[0]] = v[1]
	}
	for k, v := range args {
		m[k] = v
	}
	env = make([]string, 0, len(m))
	for k, v := range m {
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}
	return env
}

func printUsage() {
	fmt.Println(`gb-goagen, a gb plugin to run the goagen tool in a gb project.

Usage:
  gb goagen [arguments]

Either run gb-goagen from inside your desired output directory, or
specify the output directory with -o.

Example running from GB_PROJECT_DIR (see gb env):
	gb goagen bootstrap -d goa-adder/design -o $PWD/src/goa-adder

Runs as if you had executed the following:
	GOPATH=$GOPATH:$PWD goagen bootstrap -d goa-adder/design -o $PWD/src/goa-adder

Getting help:

	Sames as gb -h, shows gb help, NOT goagen help:
		gb goagen -h

	To see this help together with goagen help:
		gb goagen

	To see just the goagen help:
		goagen
		goagen -h

Output from goagen help:`)
}
