package zebra

import (
	"fmt"
	"math"
)

type constrainter interface {
	forSatisfaction() constraint
	forApplicability() constraint
}

type constraint func(...uint8) bool

func (cs constraint) forSatisfaction() constraint { return cs }
func (constraint) forApplicability() constraint   { return func(assignment ...uint8) bool { return true } }

func partialSolution(assignment ...uint8) bool {
	var count []uint8
	for i, v := range assignment {
		if i%5 == 0 {
			count = make([]uint8, 6)
		}
		count[v]++
		if v == 0 || count[v] == 1 {
			continue
		}
		return false
	}
	return true
}

func indexOf(item domainer, assignment ...uint8) (uint8, error) {
	for i := 5 * category(item); i < 5*(1+category(item)); i++ {
		if assignment[i] != item.toDomainVal() {
			continue
		}
		return uint8(i), nil
	}
	return 0, fmt.Errorf("index of %v not found inside %+v", item, assignment)
}

func otherIdx(base int, a, b domainer) int {
	delta := int(5 * math.Abs(float64(category(a)-category(b))))
	return (base + delta) % 30
}
