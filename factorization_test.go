package combicorn

import (
	"reflect"
	"testing"
)

func TestFactorMaterialization(t *testing.T) {
	factor := Factor[int]{
		Prime: 3,
		Power: 5,
	}
	mat := factor.Materialize()
	if mat != 243 {
		t.Fatalf("3**5 should be equal to 243 but results %d", mat)
	}
}

func TestFactorizedMaterialization(t *testing.T) {
	fd := Factorized[int]{
		{2, 3},
		{3, 2},
	}
	mat := fd.Materialize()
	if mat != 72 {
		t.Fatalf("(2**3)*(3**2) should be equal to 72 but results %d", mat)
	}
}

func TestFactorizedMultiplication(t *testing.T) {
	a := Factorized[int]{
		{2, 3},
		{3, 2},
		{7, 1},
	}

	b := Factorized[int]{
		{3, 1},
		{5, 2},
	}

	ab := MulFactorized(a, b).Materialize()
	abex := a.Materialize() * b.Materialize()

	if abex != ab {
		t.Fatalf("Factorized multiplication expected to give %d but gives %d", abex, ab)
	}
}

func TestFactorizedDivision(t *testing.T) {
	tdata := []struct {
		a, b   Factorized[int]
		ok     bool
		result Factorized[int]
	}{
		{Factorized[int]{{2, 3}, {3, 2}, {5, 2}}, Factorized[int]{{3, 1}}, true, Factorized[int]{{2, 3}, {3, 1}, {5, 2}}},
		{Factorized[int]{{2, 3}}, Factorized[int]{{2, 2}, {3, 1}}, false, Factorized[int]{}},
		{Factorized[int]{{3, 2}, {5, 1}}, Factorized[int]{{3, 1}, {5, 2}}, false, Factorized[int]{}},
		{Factorized[int]{{3, 2}, {7, 1}}, Factorized[int]{{3, 1}, {5, 2}}, false, Factorized[int]{}},
		{Factorized[int]{{3, 2}, {5, 5}, {7, 3}}, Factorized[int]{{3, 1}, {7, 2}, {11, 11}}, false, Factorized[int]{}},
	}

	for _, tc := range tdata {
		result, ok := DivFactorized(tc.a, tc.b)
		if ok && !tc.ok {
			t.Fatalf("Factorized division of %v by %v is expected to fail, but gives %v", tc.a, tc.b, result)
		}
		if !ok && tc.ok {
			t.Fatalf("Factorized division of %v by %v is expected to succeed, but fails", tc.a, tc.b)
		}

		if ok && !reflect.DeepEqual(result, tc.result) {
			t.Fatalf("Factorized division of %v by %v is expected to give %v, but gives %v", tc.a, tc.b, tc.result, result)
		}
	}
}
