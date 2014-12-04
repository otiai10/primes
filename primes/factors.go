package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"

	"github.com/otiai10/primes"
)

type cFactors struct {
	*cBase
	origin int64
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
	num, _ := strconv.ParseInt(m[1], 10, 64)
	c.origin = num

}

func (c *cFactors) Perform() {
	factors := primes.Factorize(c.origin)
	if c.dict {
		fmt.Println(factors.Powers())
	} else {
		fmt.Println(factors.All())
	}
}
