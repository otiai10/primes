package main

import (
	"flag"
	"fmt"
	"regexp"

	"github.com/otiai10/primes"
)

var fractionExp = regexp.MustCompile("([0-9]+)/([0-9]+)")

type cReduce struct {
	*cBase
	fraction *primes.Fraction
}

func (c *cReduce) Prepare() {
	args := flag.Args()
	if len(args) < 2 {
		c.invalid("`reduce` needs second arg like `primes reduce 144/360`.")
	}
	fraction, err := primes.ParseFractionString(args[1])
	if err != nil {
		c.invalid(err.Error())
	}
	c.fraction = fraction
}

func (c *cReduce) Perform() {
	fmt.Println(c.fraction.Reduce(-1).String())
}
