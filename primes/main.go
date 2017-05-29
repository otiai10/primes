package main

import (
	"github.com/otiai10/primes"
	"github.com/otiai10/primes/knowledge"
)

func init() {
	k, err := knowledge.NewSQLiteKnowledge(".primes")
	if err != nil {
		panic(err)
	}
	primes.Ask(k)
}

func main() {
	cmd := getCommand()
	cmd.Prepare()
	cmd.Perform()
}
