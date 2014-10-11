package srime

type Prime int

func (p Prime) canDevide(num int) bool {
	return num%int(p) == 0
}

type Primes struct {
	list []int
}

func (ps *Primes) register(num int) {
	ps.list = append(ps.list, num)
}

func (ps *Primes) canDevide(num int) bool {
	for _, prime := range ps.list {
		if Prime(prime).canDevide(num) {
			return true
		}
	}
	return false
}

func (ps *Primes) classify(num int) {
	// TODO: use goroutine to speedup
	if !ps.canDevide(num) {
		ps.register(num)
	}
}

func (ps *Primes) out() []int {
	return ps.list
}

func FindPrimesUntil(limit int) []int {
	if limit < 2 {
		return []int{}
	}
	primes := &Primes{list: []int{2}}
	if limit < 3 {
		return primes.out()
	}
	for i := 3; i < limit; i += 2 {
		primes.classify(i)
	}
	return primes.out()
}
