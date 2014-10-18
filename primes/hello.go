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

var defaultMessage = `"primes" is a simple way to handle primes and fraction with reduction.

Usage:
    primes help                        Show this message
    primes [reduce|r] <fraction> [-s]  Show reduced fraction expression     
    primes [primes|p] <max>            Show all primes until given number
    primes [factors|f] <orgin> [-p]    Show all factors of given number

Example:
    primes reduce 144/360              2/5
    primes primes 10                   [2 3 5 7]
    primes factors 924                 [2 3 7 11]
`
