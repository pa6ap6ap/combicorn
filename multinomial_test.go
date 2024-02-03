package combicorn

import (
	"testing"
)

func TestMultinomialCalculation(t *testing.T) {
	mn := Multinomial[int]{1, 2, 2, 4, 3}
	fzer := GenerateFactorsLookupTable(12)
	mnValue, _ := CalcMultinomial(fzer, mn)
	expected := 831600
	if mnValue != expected {
		t.Fatalf("Multinomial of %v is expected to equal to %d but returns %d", mn, expected, mnValue)
	}

	mn = append(mn, 1)
	mnValue, ok := CalcMultinomial(fzer, mn)
	if ok {
		t.Fatalf("Multinomial of %v is expected to fail because of exceeding upper factorization limit, but returns %d", mn, mnValue)
	}
}
