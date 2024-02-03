package combicorn

import (
	"reflect"
	"testing"
)

func TestFactorizationLookup(t *testing.T) {
	const fzlimit = 100
	fzer := GenerateFactorsLookupTable(fzlimit)

	tdata := []struct {
		n    int
		fzed Factorized[int]
		ok   bool
	}{
		{1, Factorized[int]{}, true},
		{2, Factorized[int]{{2, 1}}, true},
		{100, Factorized[int]{{2, 2}, {5, 2}}, true},
		{49, Factorized[int]{{7, 2}}, true},
		{4, Factorized[int]{{2, 2}}, true},
		{fzlimit + 1, Factorized[int]{}, false},
	}

	for _, tc := range tdata {
		fzed, ok := fzer.GetFactors(tc.n)
		if tc.ok {
			if !ok {
				t.Fatalf("Factorization of %d expected to succeed but failed", tc.n)
			}
			if !reflect.DeepEqual(fzed, tc.fzed) {
				t.Fatalf("%d should be factorized as %v but results as %v", tc.n, tc.fzed, fzed)
			}
		} else {
			if ok {
				t.Fatalf("Factorization of %d expected to fail but results as %v", tc.n, fzed)
			}
		}
	}
}
