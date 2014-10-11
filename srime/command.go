package main

import (
	"flag"
	"os"
)

type Command interface {
	Prepare()
	Perform()
}

func getCommand() Command {
	flag.Parse()
	args := getArgs()
	if len(args) == 0 {
		return &cHello{}
	}
	switch args[0] {
	case "reduce":
		return &cReduce{}
	}
	return &cHello{}
}

func getArgs() []string {
	return flag.Args()
}

func invalid(mess string) {
	os.Stderr.Write([]byte(mess))
	os.Exit(1)
}
