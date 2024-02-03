package combicorn

type Range[T Integer] struct {
	First T
	Last  T
}

type MultiRange[T Integer] []Range[T]

func SymmDiffRanges[T Integer](l, r Range[T]) (MultiRange[T], MultiRange[T]) {
	if l.First > r.Last || r.First > l.Last {
		return MultiRange[T]{l}, MultiRange[T]{r}
	}

	if r.First > l.First {
		if r.Last < l.Last {
			return MultiRange[T]{{l.First, r.First - 1}, {r.Last + 1, l.Last}}, nil
		} else if r.Last > l.Last {
			return MultiRange[T]{{l.First, r.First - 1}}, MultiRange[T]{{l.Last + 1, r.Last}}
		} else {
			return MultiRange[T]{{l.First, r.First - 1}}, nil
		}
	} else if l.First > r.First {
		if l.Last < r.Last {
			return nil, MultiRange[T]{{r.First, l.First - 1}, {l.Last + 1, r.Last}}
		} else if l.Last > r.Last {
			return MultiRange[T]{{r.Last + 1, l.Last}}, MultiRange[T]{{r.First, l.First - 1}}
		} else {
			return nil, MultiRange[T]{{r.First, l.First - 1}}
		}
	} else {
		if r.Last < l.Last {
			return MultiRange[T]{{r.Last + 1, l.Last}}, nil
		} else if l.Last < r.Last {
			return nil, MultiRange[T]{{l.Last + 1, r.Last}}
		} else {
			return nil, nil
		}
	}
}

type RangeIterator[T Integer] interface {
	Iterate(func(T) bool) bool
}

func (r Range[T]) Iterate(f func(T) bool) bool {
	for i := r.First; i <= r.Last; i++ {
		if !f(i) {
			return false
		}
	}
	return true
}

func (mr MultiRange[T]) Iterate(f func(T) bool) bool {
	for i := range mr {
		if !mr[i].Iterate(f) {
			return false
		}
	}
	return true
}

type RangeProductFraction[T Integer] struct {
	N MultiRange[T]
	D MultiRange[T]
}
