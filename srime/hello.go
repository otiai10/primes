package main

import (
	"fmt"
)

type cHello struct {
}

func (c *cHello) Prepare() {

}

func (c *cHello) Perform() {
	fmt.Println(defaultMessage)
}

var defaultMessage = `"srime" is a simple way to handle primes and fraction with reduction.

Usage:
    srime help                        Show this message
    srime [reduce|r] <fraction> [-s]  Show reduced fraction expression     
    srime [primes|p] <max>            Show all primes until given number
    srime [factors|f] <orgin> [-p]    Show all factors of given number

Example:
    srime reduce 144/360              2/5
    srime primes 10                   [2 3 5 7]
    srime factors 924                 [2 3 7 11]
`
