package main

import (
	"flag"
	"fmt"
	"github.com/otiai10/primes"
	"regexp"
	"strconv"
)

var fractionExp = regexp.MustCompile("([0-9]+)/([0-9]+)")

type cReduce struct {
	*cBase
	fraction primes.Fraction
}

func (c *cReduce) Prepare() {
	args := flag.Args()
	if len(args) < 2 {
		c.invalid("`reduce` needs second arg like `primes reduce 144/360`.")
	}
	if !fractionExp.MatchString(args[1]) {
		c.invalid("`reduce` arg must be fraction expression like `144/360`.")
	}
	m := fractionExp.FindStringSubmatch(args[1])
	nume, _ := strconv.Atoi(m[1])
	deno, _ := strconv.Atoi(m[2])
	c.fraction = primes.Fraction{nume, deno}
}

func (c *cReduce) Perform() {
	fmt.Println(c.fraction.Reduce().String())
}
