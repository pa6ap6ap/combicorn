package combicorn

type Multinomial[T Integer] []T

type MultinomialCalculator[T Integer] interface {
	Calculate(Multinomial[T]) T
}

type lookupBasedMC[T Integer] FactorsLookupTable[T]

func MultinomialToRangeProductFraction[T Integer](m Multinomial[T]) (nom Range[T], denom MultiRange[T]) {
	reducedComponent := -1
	denomLen := -1
	for i, mi := range m {
		nom.Last += mi
		if reducedComponent == -1 || m[reducedComponent] < mi {
			reducedComponent = i
		}
		if mi > 1 {
			denomLen++
		}
	}
	denom = make(MultiRange[T], 0, denomLen)
	for i, mi := range m {
		if i != reducedComponent && mi > 1 {
			denom = append(denom, Range[T]{2, mi})
		}
	}
	nom.First = m[reducedComponent] + 1
	return
}

func FactorizeProduct[T Integer](fzer Factorizer[T], rg RangeIterator[T]) (fzed Factorized[T], ok bool) {
	ok = rg.Iterate(func(i T) bool {
		factors, ok := fzer.GetFactors(i)
		if !ok {
			return false
		}
		fzed = MulFactorized(fzed, factors)
		return true
	})
	return
}

func CalcMultinomial[T Integer](fzer Factorizer[T], mn Multinomial[T]) (T, bool) {
	nom, denom := MultinomialToRangeProductFraction(mn)
	nomFz, ok := FactorizeProduct[T](fzer, nom)
	if !ok {
		return 0, false
	}
	denomFz, _ := FactorizeProduct[T](fzer, denom)
	fzed, _ := DivFactorized[T](nomFz, denomFz)
	return fzed.Materialize(), ok
}
