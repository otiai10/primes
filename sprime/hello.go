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

var defaultMessage = `"sprime" is a simple way to handle primes and fraction with reduction.

Usage:
    sprime help                        Show this message
    sprime [reduce|r] <fraction> [-s]  Show reduced fraction expression     
    sprime [primes|p] <max>            Show all primes until given number
    sprime [factors|f] <orgin> [-p]    Show all factors of given number

Example:
    sprime reduce 144/360              2/5
    sprime primes 10                   [2 3 5 7]
    sprime factors 924                 [2 3 7 11]
`
