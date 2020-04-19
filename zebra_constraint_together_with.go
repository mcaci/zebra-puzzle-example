package zebra

type togetherWithConstraint struct {
	a, b domainer
}

type togetherWithConstraintData struct {
	otherIdx      int
	oneIndexFound bool
}

func (twc togetherWithConstraint) dataFrom(assignment ...uint8) togetherWithConstraintData {
	aIdx, aErr := indexOf(twc.a, assignment...)
	bIdx, bErr := indexOf(twc.b, assignment...)
	oneIndexFound := aErr == nil || bErr == nil
	switch {
	case !oneIndexFound:
		return togetherWithConstraintData{oneIndexFound: oneIndexFound}
	case aErr == nil:
		return togetherWithConstraintData{oneIndexFound: oneIndexFound, otherIdx: otherIdx(int(aIdx), twc.a, twc.b)}
	case bErr == nil:
		return togetherWithConstraintData{oneIndexFound: oneIndexFound, otherIdx: otherIdx(int(bIdx), twc.a, twc.b)}
	default:
		return togetherWithConstraintData{oneIndexFound: oneIndexFound}
	}
}

func (twc togetherWithConstraint) forApplicability() constraint {
	return func(assignment ...uint8) bool {
		data := twc.dataFrom(assignment...)
		return data.oneIndexFound && assignment[data.otherIdx] != 0
	}
}

func (twc togetherWithConstraint) forSatisfaction() constraint {
	return func(assignment ...uint8) bool {
		aIdx, aErr := indexOf(twc.a, assignment...)
		bIdx, bErr := indexOf(twc.b, assignment...)
		return aErr == nil && bErr == nil && (bIdx%5 == aIdx%5)
	}
}

func (twc togetherWithConstraint) forMissingIndex(assignment []uint8) (uint8, bool) {
	data := twc.dataFrom(assignment...)
	return uint8(data.otherIdx), data.oneIndexFound
}
