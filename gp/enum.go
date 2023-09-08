package gp

type DirectionSign int16

const (
	DirectionSignNone DirectionSign = iota
	Coda
	DoubleCoda
	Segno
	SegnoSegno
	Fine
	DaCapo
	DaCapoAlCoda
	DaCapoAlDoubleCoda
	DaCapoAlFine
	DaSegno
	DaSegnoAlCoda
	DaSegnoAlDoubleCoda
	DaSegnoAlFine
	DaSegnoSegno
	DaSegnoSegnoAlCoda
	DaSegnoSegnoAlDoubleCoda
	DaSegnoSegnoAlFine
	DaCoda
	DaDoubleCoda
)

type TripletFeel uint8

const (
	TripletFeelNone      TripletFeel = 0
	TripletFeelEighth                = 1
	TripletFeelSixteenth             = 2
)

var toTripletFeel = map[int8]TripletFeel{
	0: TripletFeelNone,
	1: TripletFeelEighth,
	2: TripletFeelSixteenth,
}
