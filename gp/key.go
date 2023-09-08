package gp

const (
	DURATION_QUARTER_TIME          int64 = 960
	DURATION_QUARTER               uint8 = 4
	DURATION_EIGHTH                uint8 = 8
	DURATION_SIXTEENTH             uint8 = 16
	DURATION_THIRTY_SECOND         uint8 = 32
	DURATION_SIXTY_FOURTH          uint8 = 64
	DURATION_HUNDRED_TWENTY_EIGHTH uint8 = 128
)

type KeySignature struct {
	Key     int8
	IsMinor bool
}

type TimeSignature struct {
	Numerator   int8
	Denominator Duration
	Beams       []uint8
}

var DefaultTimeSignature = TimeSignature{
	Numerator:   4,
	Denominator: DefaultDuration,
	Beams:       []uint8{2, 2, 2, 2},
}

type Duration struct {
	Value        uint16
	Dotted       bool
	DoubleDotted bool
	MinTime      uint8
	TupletEnters uint8
	TupletTimes  uint8
}

var DefaultDuration = Duration{
	Value:        uint16(DURATION_QUARTER),
	TupletEnters: 1,
	TupletTimes:  1,
}
