package srime

type Reduced struct {
	origin Fraction
	result Fraction
}

func Reduce(numerator, denomirator int) *Reduced {
	reduced := &Reduced{
		Fraction{numerator, denomirator},
		Fraction{numerator, denomirator}.Reduce(),
	}
	return reduced
}

func (r *Reduced) String() string {
	return r.result.String()
}
