package primes

type cache struct {
	store map[int64]*Primes
}

var globally = &cache{
	store: map[int64]*Primes{},
}

func (c *cache) Know(target int64) *Primes {
	if p, ok := c.store[target]; ok {
		return p
	}
	return nil
}

func (c *cache) Learn(primes *Primes) error {
	c.store[primes.target] = primes
	return nil
}
