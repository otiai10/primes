package primes

import "container/list"

// Knowledge ...
type Knowledge interface {
	Init() error
	Clear() error
	Know(int64) *Primes
	Learn(*Primes) Knowledge
	Persist(int64) *Primes
	Until(int64) *Primes
}

func extends(base *Primes, target int64) *Primes {
	p := new(Primes)
	p.target = target
	p.dictionary = base.dictionary
	p.list = base.list
	return p
}

var knowledge Knowledge = &OnMemoryKnowledge{
	store: map[int64]*Primes{},
	list:  list.New(),
}

// UseKnowledge ...
func UseKnowledge(k Knowledge) error {
	knowledge = k
	return knowledge.Init()
}
