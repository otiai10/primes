package main

import (
	"flag"
	"log"
)

// Command interface
type Command interface {
	Prepare()
	Perform()
}

func getCommand() Command {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		return &cHello{}
	}
	switch args[0] {
	case "reduce", "r":
		return &cReduce{}
	case "factors", "f":
		return &cFactors{}
	case "primes", "p":
		return &cPrimes{}
	}
	return &cHello{}
}

type cBase struct{}

func (c *cBase) invalid(mess string) {
	log.Fatalln(mess)
}
