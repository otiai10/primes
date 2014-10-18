package main

import (
	"flag"
	"fmt"
	"github.com/otiai10/primes"
	"regexp"
	"strconv"
	"strings"
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

	flag.BoolVar(&c.dict, "dict", true, "-dict options shows also times of factors")
}

func (c *cFactors) Perform() {
	factors := primes.Factorize(float64(c.origin))
	if c.dict {
		fmt.Println(c.expressDict(factors.Dict()))
	} else {
		fmt.Println(c.expressList(factors.List()))
	}
}

func (c *cFactors) expressDict(dict map[int]int) string {
	var pool = []string{}
	for i, v := range dict {
		pool = append(pool, fmt.Sprintf("%d^%d", i, v))
	}
	return "[" + strings.Join(pool, " ") + "]"
}
func (c *cFactors) expressList(list []int) string {
	return fmt.Sprintln(list)
}
