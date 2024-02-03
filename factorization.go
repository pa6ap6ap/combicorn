package combicorn

type Factor[T Integer] struct {
	Prime T
	Power T
}

func (fc Factor[T]) Materialize() T {
	var acc T = 1
	p := fc.Prime
	for fc.Power != 0 {
		if fc.Power&1 != 0 {
			acc *= p
		}
		fc.Power >>= 1
		p *= p
	}
	return acc
}

type Factorized[T Integer] []Factor[T]

func (fd Factorized[T]) Materialize() T {
	var acc T = 1
	for _, fc := range fd {
		acc *= fc.Materialize()
	}
	return acc
}

func MulFactorized[T Integer](a, b Factorized[T]) Factorized[T] {
	prod := make(Factorized[T], 0, len(a)+len(b))
	for len(a) != 0 && len(b) != 0 {
		if a[0].Prime < b[0].Prime {
			prod = append(prod, a[0])
			a = a[1:]
		} else if b[0].Prime < a[0].Prime {
			prod = append(prod, b[0])
			b = b[1:]
		} else {
			prod = append(prod, Factor[T]{a[0].Prime, a[0].Power + b[0].Power})
			a, b = a[1:], b[1:]
		}
	}
	prod = append(append(prod, a...), b...)

	return prod
}

func DivFactorized[T Integer](a, b Factorized[T]) (Factorized[T], bool) {
	if len(a) < len(b) {
		return nil, false
	}
	res := make(Factorized[T], 0, len(a))
	for len(a) != 0 && len(b) != 0 {
		if a[0].Prime < b[0].Prime {
			res, a = append(res, a[0]), a[1:]
		} else if a[0].Prime == b[0].Prime {
			if b[0].Power < a[0].Power {
				res = append(res, Factor[T]{a[0].Prime, a[0].Power - b[0].Power})
			} else if a[0].Power < b[0].Power {
				return nil, false
			}
			a, b = a[1:], b[1:]
		} else {
			return nil, false
		}
	}

	if len(b) != 0 {
		return nil, false
	}
	res = append(res, a...)

	return res, true
}
