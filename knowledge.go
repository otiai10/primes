package primes

import "container/list"

// Knowledge ...
type Knowledge interface {
	Init() error
	Clear() error
	Know(int64) *Primes
	Learn(*Primes) Knowledge
	Until(int64) *Primes
}

// OnMemoryKnowledge ...
type OnMemoryKnowledge struct {
	store map[int64]*Primes
	list  *list.List
}

// Init ...
func (m *OnMemoryKnowledge) Init() error {
	m.store = map[int64]*Primes{}
	m.list = list.New()
	return nil
}

// Clear ...
func (m *OnMemoryKnowledge) Clear() error {
	return m.Init()
}

// Know for interface `Knowledge`.
func (m *OnMemoryKnowledge) Know(n int64) *Primes {
	if p, ok := m.store[n]; ok {
		return p
	}
	return nil
}

// Learn for interface `Knowledge`.
func (m *OnMemoryKnowledge) Learn(p *Primes) Knowledge {
	m.store[p.target] = p
	if m.list.Len() == 0 {
		m.list.PushFront(p)
		return m
	}
	for e := m.list.Front(); e != nil; e = e.Next() {
		if e.Value.(*Primes).target > p.target {
			m.list.InsertBefore(p, e)
			return m
		}
	}
	m.list.PushBack(p)
	return m
}

// Until for interface `Knowledge`.
func (m *OnMemoryKnowledge) Until(n int64) *Primes {
	base := m.Persist(n)

	i := base.target
	if i == 1 {
		i = 2
	}
	p := extends(base, n)

	for ; i <= p.target; i++ {
		if p.knows(i) {
			continue // needless to evaluate.
		}
		if p.canDevideByKnownPrimeNumbers(i) {
			continue // it's not prime number.
		}
		// it's prime number,
		// and multiples of this number are no longer needless to be eveluated
		p.add(i)
	}

	m.Learn(p)

	return p
}

// Persist persists this knowledge already knows everything until target.
func (m *OnMemoryKnowledge) Persist(n int64) *Primes {
	if m.list.Len() == 0 {
		return Until(2)
	}
	max := m.list.Back().Value.(*Primes)
	fabricated := &Primes{target: n, dictionary: map[int64]bool{}, list: []int64{}}
	for _, i := range max.List() {
		if i < n {
			fabricated.add(i)
		}
	}
	return fabricated
}

func clone(base *Primes, target int64) *Primes {
	p := new(Primes)
	p.target = target
	p.dictionary = base.dictionary
	p.list = base.list
	return p
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
