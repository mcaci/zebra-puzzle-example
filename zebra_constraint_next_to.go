package zebra

import (
	"math"
)

type nextToConstraint struct {
	a, b domainer
}

type nextToConstraintData struct {
	aIdx, bIdx, aHouseNumIdx, bHouseNumIdx int
	indexesFound                           bool
}

func (ntc nextToConstraint) dataFrom(assignment ...uint8) nextToConstraintData {
	aIdx, aErr := indexOf(ntc.a, assignment...)
	bIdx, bErr := indexOf(ntc.b, assignment...)
	indexesFound := aErr == nil && bErr == nil
	nIdx, mIdx := otherIdx(int(aIdx), ntc.a, number(0)), otherIdx(int(bIdx), ntc.b, number(0))
	return nextToConstraintData{aIdx: int(aIdx), bIdx: int(bIdx), indexesFound: indexesFound, aHouseNumIdx: nIdx, bHouseNumIdx: mIdx}
}

func (ntc nextToConstraint) forApplicability() constraint {
	return func(assignment ...uint8) bool {
		data := ntc.dataFrom(assignment...)
		return data.indexesFound && assignment[data.aHouseNumIdx] != 0 && assignment[data.bHouseNumIdx] != 0
	}
}

func (ntc nextToConstraint) forSatisfaction() constraint {
	return func(assignment ...uint8) bool {
		data := ntc.dataFrom(assignment...)
		return math.Abs(float64(assignment[data.aHouseNumIdx])-float64(assignment[data.bHouseNumIdx])) == 1 &&
			(assignment[data.aIdx]%5 != assignment[data.bIdx]%5)
	}
}

func (ntc nextToConstraint) forMissingIndex(assignment []uint8) (uint8, bool) {
	data := ntc.dataFrom(assignment...)
	switch {
	case !data.indexesFound:
		return 0, false
	case assignment[data.aHouseNumIdx] == 0:
		return uint8(data.aHouseNumIdx), true
	case assignment[data.bHouseNumIdx] == 0:
		return uint8(data.bHouseNumIdx), true
	default:
		return 0, false
	}
}

type toTheRightOfConstraint nextToConstraint

func (ttroc toTheRightOfConstraint) forApplicability() constraint {
	return nextToConstraint(ttroc).forApplicability()
}

func (ttroc toTheRightOfConstraint) forSatisfaction() constraint {
	return func(assignment ...uint8) bool {
		data := nextToConstraint(ttroc).dataFrom(assignment...)
		return (assignment[data.aHouseNumIdx]-assignment[data.bHouseNumIdx] == 1) &&
			(assignment[data.aIdx]%5 != assignment[data.bIdx]%5)
	}
}
