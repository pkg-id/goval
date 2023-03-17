package goval

type RuleCode int

const (
	ruleCodeNilBase RuleCode = (1 + iota) * 1_000
	ruleCodeBaseString
	ruleCodeBaseNumber
	ruleCodeBaseSlice
	ruleCodeBaseMap
)

const (
	NilRequired = ruleCodeNilBase + iota
)

const (
	StringRequired = ruleCodeBaseString + iota
	StringMin
	StringMax
	StringMatch
)

const (
	NumberRequired = ruleCodeBaseNumber + iota
	NumberMin
	NumberMax
)

const (
	SliceRequired = ruleCodeBaseSlice + iota
	SliceMin
	SliceMax
	SliceEach
)

const (
	MapRequired = ruleCodeBaseMap + iota
	MapMin
	MapMax
	MapEach
)

func IsCodeEqual(a, b RuleCode) bool {
	return a == b
}
