package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/otiai10/primes"
)

type cPrimes struct {
	*cBase
	limit int64
}

func (c *cPrimes) Prepare() {
	args := flag.Args()
	if len(args) < 2 {
		c.invalid("`primes` needs second arg like `primes primes 12`.")
	}
	if !numericExp.MatchString(args[1]) {
		c.invalid("`primes` arg must be number expression like `12345`.")
	}
	m := numericExp.FindStringSubmatch(args[1])
	num, _ := strconv.ParseInt(m[1], 10, 64)
	c.limit = num
}

func (c *cPrimes) Perform() {
	fmt.Println(primes.Until(c.limit).List())
}
