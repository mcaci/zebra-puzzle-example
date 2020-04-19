package zebra

// CSP is a structure that models a
// simple constraint satisfaction problem
type CSP struct {
	v, d []uint8
	c    []constrainter
}

// SolvePuzzle function solves the zebra puzzle
func SolvePuzzle() Solution {
	csp := CSP{
		v: make([]uint8, 30),
		d: []uint8{1, 2, 3, 4, 5},
		c: []constrainter{
			constraint(partialSolution),
			togetherWithConstraint{a: englishman, b: red},
			togetherWithConstraint{a: green, b: coffee},
			togetherWithConstraint{a: spaniard, b: dog},
			togetherWithConstraint{a: ukrainian, b: tea},
			toTheRightOfConstraint{a: green, b: ivory},
			togetherWithConstraint{a: oldGold, b: snails},
			togetherWithConstraint{a: kools, b: yellow},
			togetherWithConstraint{a: milk, b: house3},
			togetherWithConstraint{a: norwegian, b: house1},
			nextToConstraint{a: chesterfields, b: fox},
			nextToConstraint{a: kools, b: horse},
			togetherWithConstraint{a: luckyStrike, b: orangeJuice},
			togetherWithConstraint{a: japanese, b: parliaments},
			nextToConstraint{a: norwegian, b: blue},
		},
	}
	recursiveBacktracking(&csp.v, &csp)
	return Solution{OwnsZebra: ownerOf(zebra, csp.v...), DrinksWater: ownerOf(water, csp.v...)}
}

func recursiveBacktracking(assignment *[]uint8, csp *CSP) bool {
	if complete(csp.v) {
		return true
	}
	idx := selectUnassigned(&csp.v, assignment, *csp)
domainValueLoop:
	for _, d := range csp.d {
		csp.v[idx] = d
		for _, c := range csp.c {
			switch {
			case !c.forApplicability()(csp.v...):
				continue
			case !c.forSatisfaction()(csp.v...):
				csp.v[idx] = 0
				continue domainValueLoop
			default:
			}
		}
		outcome := recursiveBacktracking(assignment, csp)
		if !outcome {
			csp.v[idx] = 0
			continue domainValueLoop
		}
		return true
	}
	return false
}

func selectUnassigned(variables, assignment *[]uint8, csp CSP) uint8 {
	for _, c := range csp.c {
		indexer := func([]uint8) (uint8, bool) { return 0, false }
		switch constr := c.(type) {
		case togetherWithConstraint:
			indexer = constr.forMissingIndex
		case nextToConstraint:
			indexer = constr.forMissingIndex
		}
		idx, ok := indexer(*variables)
		if !ok || (*variables)[idx] != 0 {
			continue
		}
		return uint8(idx)
	}
	for i, v := range *assignment {
		if v != 0 {
			continue
		}
		return uint8(i)
	}
	return 0
}

func complete(assignment []uint8) bool {
	for _, v := range assignment {
		if v != 0 {
			continue
		}
		return false
	}
	return true
}

// Solution type describes the solution of the zebra puzzle
type Solution struct {
	DrinksWater string
	OwnsZebra   string
}
