package main

import (
	"fmt"
	"github.com/otiai10/primes"
	"strconv"
)

type cPrimes struct {
	limit int
}

func (c *cPrimes) Prepare() {
	args := getArgs()
	if len(args) < 2 {
		invalid("`primes` needs second arg like `primes primes 12`.")
		return
	}
	if !numericExp.MatchString(args[1]) {
		invalid("`primes` arg must be number expression like `12345`.")
		return
	}
	m := numericExp.FindStringSubmatch(args[1])
	num, _ := strconv.Atoi(m[1])
	c.limit = num
}

func (c *cPrimes) Perform() {
	fmt.Println(primes.FindPrimesUntil(c.limit))
}
