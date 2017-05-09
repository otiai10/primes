package primes

// Knowledge interface provides the way to access
// known prime numbers, which are obtained so far.
type Knowledge interface {
	Open() error
	Close()
	Know(int64) *Primes
	Learn(*Primes)
}

var knowledge Knowledge = &EmptyKnowledge{}

// Ask asks help to Knowledge.
func Ask(k Knowledge) {
	knowledge = k
}
