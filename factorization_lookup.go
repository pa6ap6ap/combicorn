package combicorn

type Factorizer[T Integer] interface {
	GetFactors(T) (Factorized[T], bool)
}

type FactorsLookupTable[T Integer] []Factorized[T]

func (lut FactorsLookupTable[T]) GetFactors(x T) (Factorized[T], bool) {
	var factor2 T
	for x&1 == 0 {
		x >>= 1
		factor2++
	}
	i := (x / 2) - 1
	if i < 0 {
		if factor2 > 0 {
			return Factorized[T]{{2, factor2}}, true
		} else {
			return Factorized[T]{}, true
		}
	} else if i >= T(len(lut)) {
		return nil, false
	}
	if factor2 == 0 {
		return lut[i], true
	} else {
		return append(Factorized[T]{{2, factor2}}, lut[i]...), true
	}
}

func GenerateFactorsLookupTable[T Integer](limit T) Factorizer[T] {
	lutSize := (limit - 1) / 2
	lut := make(FactorsLookupTable[T], lutSize)
	appendFactor := func(x, prime T) {
		i := x/2 - 1
		if len(lut[i]) > 0 && lut[i][len(lut[i])-1].Prime == prime {
			lut[i][len(lut[i])-1].Power++
		} else {
			lut[i] = append(lut[i], Factor[T]{prime, 1})
		}
	}
	for i := T(0); i < lutSize; i++ {
		if len(lut[i]) == 0 {
			x := i*2 + 3
			for y := x; y <= limit; y *= x {
				for p := y; p <= limit; p += 2 * y {
					appendFactor(p, x)
				}
			}
		}
	}
	return lut
}
