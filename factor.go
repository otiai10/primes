package srime

type Factors struct {
	of   float64
	list []int
	dict map[int]int
}

func Factorize(num float64) *Factors {
	factors := &Factors{of: num, list: []int{}, dict: map[int]int{}}
	for _, prime := range FindPrimesUntil(int(num/2) + 1) {
		factors.classify(prime)
	}
	return factors
}

func (f *Factors) classify(possiblePrime int) {
	if Prime(possiblePrime).canDevide(int(f.of)) {
		f.register(possiblePrime)
	}
}

func (f *Factors) register(factor int) {
	f.list = append(f.list, factor)
	f.dict[factor] = factor
}

func (f *Factors) List() []int {
	return f.list
}

func (f *Factors) Dict() map[int]int {
	return f.dict
}
