package main

import (
	"fmt"
	"github.com/otiai10/srime"
	"regexp"
	"strconv"
)

var fractionExp = regexp.MustCompile("([0-9]+)/([0-9]+)")

type cReduce struct {
	fraction srime.Fraction
}

func (c *cReduce) Prepare() {
	args := getArgs()
	if len(args) < 2 {
		invalid("`reduce` needs second arg like `srime reduce 144/360`.")
		return
	}
	if !fractionExp.MatchString(args[1]) {
		invalid("`reduce` arg must be fraction expression like `144/360`.")
		return
	}
	m := fractionExp.FindStringSubmatch(args[1])
	nume, _ := strconv.Atoi(m[1])
	deno, _ := strconv.Atoi(m[2])
	c.fraction = srime.Fraction{nume, deno}
}

func (c *cReduce) Perform() {
	fmt.Println(c.fraction.Reduce().String())
}
