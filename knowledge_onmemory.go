package primes

import "container/list"

// Registry is an instance of on-memory knowledge.
type Registry struct {
	store map[int64]*Primes
	list  *list.List
}

// NewOnMemoryKnowledge ...
func NewOnMemoryKnowledge() *Registry {
	return &Registry{
		store: map[int64]*Primes{},
		list:  list.New(),
	}
}

// Open ...
func (r *Registry) Open() error {
	return nil
}

// Close ...
func (r *Registry) Close() {
	// do nothing
}

// Know returns prime number if it knows that.
func (r *Registry) Know(target int64) *Primes {
	if p, ok := r.store[target]; ok {
		return p
	}
	return nil
}

// Learn adds known prime number to registry.
func (r *Registry) Learn(p *Primes) {

	// Register it to dictionary.
	r.store[p.Target] = p

	// Just push it if its list is empty.
	if r.list.Len() == 0 {
		r.list.PushFront(p)
		return
	}

	// If there is a bigger target, insert it before that.
	for e := r.list.Front(); e != nil; e = e.Next() {
		if e.Value.(*Primes).Target > p.Target {
			r.list.InsertBefore(p, e)
			return
		}
	}

	// If it's the largest one, push to the last.
	r.list.PushBack(p)
	return
}
