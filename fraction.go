package primes

import (
	"math"
	"strconv"
)

type Fraction struct {
	Numerator   int
	Denomirator int
}

func (f Fraction) Reduce() Fraction {
	for _, common := range f.FindCommonFactors() {
		f.Numerator = f.Numerator / common
		f.Denomirator = f.Denomirator / common
	}
	return f
}

func (f Fraction) FindCommonFactors() (commons []int) {
	factorsOfN := Factorize(float64(f.Numerator))
	factorsOfD := Factorize(float64(f.Denomirator))
	for _, factor := range factorsOfN.List() {
		if factorsOfD.Has(factor) {
			pow := f.countCommonPowers(factorsOfN, factorsOfD, factor)
			commons = append(commons, int(math.Pow(float64(factor), pow)))
		}
	}
	return commons
}

// FIXME: better way?
func (f Fraction) countCommonPowers(f1, f2 *Factors, factor int) float64 {
	count1 := f1.Dict()[factor]
	count2 := f2.Dict()[factor]
	if count1 < count2 {
		return float64(count1)
	}
	return float64(count2)
}

func (f Fraction) String() string {
	return strconv.Itoa(f.Numerator) + "/" + strconv.Itoa(f.Denomirator)
}
