package main

import (
	"flag"
	"fmt"
	"github.com/otiai10/primes"
	"regexp"
	"strconv"
)

type cFactors struct {
	*cBase
	origin int
	dict   bool
}

var numericExp = regexp.MustCompile("([0-9]+)")

func (c *cFactors) Prepare() {
	args := flag.Args()
	if len(args) < 2 {
		c.invalid("`factors` needs second arg like `primes factors 12`.")
		return
	}
	if !numericExp.MatchString(args[1]) {
		c.invalid("`factors` arg must be number expression like `12345`.")
		return
	}
	m := numericExp.FindStringSubmatch(args[1])
	num, _ := strconv.Atoi(m[1])
	c.origin = num
}

func (c *cFactors) Perform() {
	fmt.Println(primes.Factorize(float64(c.origin)).List())
}
