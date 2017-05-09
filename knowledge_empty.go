package primes

// EmptyKnowledge ...
type EmptyKnowledge struct{}

// Open ...
func (e *EmptyKnowledge) Open() error {
	return nil
}

// Close ...
func (e *EmptyKnowledge) Close() {}

// Know ...
func (e *EmptyKnowledge) Know(target int64) *Primes {
	return nil
}

// Learn ...
func (e *EmptyKnowledge) Learn(p *Primes) {

}
