package combicorn

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMultiRangeIterator(t *testing.T) {
	mr := MultiRange[int]{{3, 6}, {1, 4}}
	result := ""
	mr.Iterate(func(i int) bool {
		result += fmt.Sprint(i)
		if i == 2 {
			return false
		}
		return true
	})
	if expected := "345612"; result != expected {
		t.Fatalf("Iteration result should be %s but instead is %s", expected, result)
	}
}

func TestRangesSymmetricDifference(t *testing.T) {
	type rg = Range[int]
	type mrg = MultiRange[int]

	tdata := []struct {
		l, r     rg
		lex, rex mrg
	}{
		{rg{2, 3}, rg{4, 5}, mrg{{2, 3}}, mrg{{4, 5}}},
		{rg{2, 4}, rg{4, 5}, mrg{{2, 3}}, mrg{{5, 5}}},
		{rg{2, 6}, rg{4, 5}, mrg{{2, 3}, {6, 6}}, nil},
		{rg{2, 6}, rg{4, 6}, mrg{{2, 3}}, nil},
		{rg{4, 5}, rg{2, 3}, mrg{{4, 5}}, mrg{{2, 3}}},
		{rg{4, 5}, rg{2, 4}, mrg{{5, 5}}, mrg{{2, 3}}},
		{rg{4, 5}, rg{2, 6}, nil, mrg{{2, 3}, {6, 6}}},
		{rg{4, 6}, rg{2, 6}, nil, mrg{{2, 3}}},
		{rg{4, 7}, rg{4, 7}, nil, nil},
		{rg{4, 7}, rg{4, 6}, mrg{{7, 7}}, nil},
		{rg{4, 6}, rg{4, 7}, nil, mrg{{7, 7}}},
	}
	for _, tc := range tdata {
		lr, rr := SymmDiffRanges(tc.l, tc.r)
		if !reflect.DeepEqual(lr, tc.lex) || !reflect.DeepEqual(rr, tc.rex) {
			t.Fatalf("Symmetric difference of %v and %v expected to produce %v and %v, but produced %v and %v", tc.l, tc.r, tc.lex, tc.rex, lr, rr)
		}
	}
}
