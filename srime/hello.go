package main

import (
	"github.com/otiai10/srime"
)

type cHello struct {
}

func (c *cHello) Prepare() {

}

func (c *cHello) Perform() {
	srime.Hello()
}
