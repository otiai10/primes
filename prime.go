package primes

// Primes represents prime numbers
type Primes struct {
	target int64
	field  map[int64]bool
	list   []int64
}

// Until finds
func Until(n int64) *Primes {
	p := new(Primes)
	p.target = n
	p.field = map[int64]bool{}

	for i := 2; int64(i) <= p.target; i++ {
		p.field[int64(i)] = false
	}

	for i := 2; int64(i) <= p.target; i++ {
		j := int64(i)
		if p.has(j) {
			continue
		}
		if p.canDevide(j) {
			continue
		}
		p.add(j)
	}
	return p
}

func (p *Primes) canDevide(i int64) bool {
	for _, f := range p.list {
		if i%f == 0 {
			return true
		}
	}
	return false
}

func (p *Primes) has(i int64) bool {
	marked, ok := p.field[i]
	if !ok {
		return false
	}
	return marked
}

func (p *Primes) add(i int64) {
	// register this number
	p.list = append(p.list, i)
	// mark this number
	p.field[i] = true
	// mark multiples of this number
	for j := 2; i*int64(j) < p.target; j++ {
		p.field[i*int64(j)] = true
	}
}

// List returns all found primes
func (p *Primes) List() []int64 {
	return p.list
}
