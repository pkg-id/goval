package goval

type ruleCode int

func (r ruleCode) Equal(other RuleCoder) bool {
	v, ok := other.(ruleCode)
	return ok && r == v
}

type RuleCoder interface {
	Equal(other RuleCoder) bool
}

const (
	rcPointer ruleCode = (1 + iota) * 1_000
	rcString
	rcNumber
	rcSlice
	rcMap
	rcTime
)

const (
	PtrRequired = rcPointer + iota
)

const (
	StringRequired = rcString + iota
	StringMin
	StringMax
	StringMatch
)

const (
	NumberRequired = rcNumber + iota
	NumberMin
	NumberMax
)

const (
	SliceRequired = rcSlice + iota
	SliceMin
	SliceMax
	SliceEach
)

const (
	MapRequired = rcMap + iota
	MapMin
	MapMax
	MapEach
)

const (
	TimeRequired = rcTime + iota
	TimeMin
	TimeMax
)
