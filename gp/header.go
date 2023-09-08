package gp

type Version struct {
	Data      string
	Number    [3]uint
	Clipboard bool
}

type Clipboard struct {
	// TODO
}

type MeasureHeader struct {
	Number            uint16
	Start             int64
	Tempo             int32
	DoubleBar         bool
	RepeatOpen        bool
	RepeatAlternative uint8
	RepeatClose       int8
	TripletFeel       TripletFeel
	Direction         DirectionSign
	KeySignature      KeySignature
	TimeSignature     TimeSignature
	Marker            *Marker
}

var DefaultMeasureHeader = MeasureHeader{
	Number:        1,
	Start:         DURATION_QUARTER_TIME,
	RepeatClose:   -1,
	TimeSignature: DefaultTimeSignature,
}

type Marker struct {
	Title string
	Color int32
}

var DefaultMarker = Marker{
	Title: "Section",
	Color: 0xff0000,
}

// TODO
func readMarker(data []byte, off *uint) *Marker {
	m := DefaultMarker
	m.Title = readIntSizeString(data, off)
	m.Color = readColor(data, off)
	return &m
}

func (s *Song) readClipboard(data []byte, off *uint) Clipboard {
	var c Clipboard
	if !s.Version.Clipboard {
		return c
	}
	return c
}

func (s *Song) readDirections(data []byte, off *uint) (map[DirectionSign]int16, map[DirectionSign]int16) {
	signs := make(map[DirectionSign]int16, 5)
	signs[Coda] = readShort(data, off)
	signs[DoubleCoda] = readShort(data, off)
	signs[Segno] = readShort(data, off)
	signs[SegnoSegno] = readShort(data, off)
	signs[Fine] = readShort(data, off)

	fromSigns := make(map[DirectionSign]int16, 14)
	fromSigns[DaCapo] = readShort(data, off)
	fromSigns[DaCapoAlCoda] = readShort(data, off)
	fromSigns[DaCapoAlDoubleCoda] = readShort(data, off)
	fromSigns[DaCapoAlFine] = readShort(data, off)
	fromSigns[DaSegno] = readShort(data, off)
	fromSigns[DaSegnoAlCoda] = readShort(data, off)
	fromSigns[DaSegnoAlDoubleCoda] = readShort(data, off)
	fromSigns[DaSegnoAlFine] = readShort(data, off)
	fromSigns[DaSegnoSegno] = readShort(data, off)
	fromSigns[DaSegnoSegnoAlCoda] = readShort(data, off)
	fromSigns[DaSegnoSegnoAlDoubleCoda] = readShort(data, off)
	fromSigns[DaSegnoSegnoAlFine] = readShort(data, off)
	fromSigns[DaCoda] = readShort(data, off)
	fromSigns[DaDoubleCoda] = readShort(data, off)

	return signs, fromSigns
}

func (s *Song) readMeasureHeadersV5(data []byte, off *uint, measureCount uint, signs, fromSigns map[DirectionSign]int16) {
	var prev *MeasureHeader
	for i := uint(1); i < measureCount+1; i++ {
		mh, _ := s.readMeasureHeaderV5(data, off, i, prev)
		prev = mh
		s.MeasureHeaders = append(s.MeasureHeaders, *mh)
	}
	for k, v := range signs {
		if v > -1 {
			s.MeasureHeaders[v-1].Direction = k
		}
	}
	for k, v := range fromSigns {
		if v > -1 {
			s.MeasureHeaders[v-1].Direction = k
		}
	}
}

func (s *Song) readMeasureHeader(data []byte, off *uint, num uint, prev *MeasureHeader) (*MeasureHeader, uint8) {
	flag := readByte(data, off)
	mh := DefaultMeasureHeader
	mh.Number = uint16(num)
	mh.Start = 0
	mh.TripletFeel = s.TripletFeel
	if (flag & 0x01) == 0x01 {
		mh.TimeSignature.Numerator = readSignedByte(data, off)
	} else if num > 1 {
		mh.TimeSignature.Numerator = prev.TimeSignature.Numerator
	}

	if (flag & 0x02) == 0x02 {
		mh.TimeSignature.Denominator.Value = uint16(readSignedByte(data, off))
	} else if num > 1 {
		mh.TimeSignature.Denominator.Value = prev.TimeSignature.Denominator.Value
	}

	if (flag & 0x08) == 0x08 {
		mh.RepeatClose = readSignedByte(data, off)
	}

	if (flag & 0x10) == 0x10 {
		if s.Version.Number[0] == 5 {
			mh.RepeatAlternative = s.readRepeatAlternativeV5(data, off)
		} else {
			mh.RepeatAlternative = s.readRepeatAlternative(data, off)
		}
	}

	if (flag & 0x20) == 0x20 {
		mh.Marker = readMarker(data, off)
	}

	if (flag & 0x40) == 0x40 {
		mh.KeySignature.Key = readSignedByte(data, off)
		mh.KeySignature.IsMinor = readSignedByte(data, off) != 0
	} else if mh.Number > 1 {
		mh.KeySignature = prev.KeySignature
	}
	mh.DoubleBar = (flag & 0x80) == 0x80
	return &mh, flag
}

func (s *Song) readMeasureHeaderV5(data []byte, off *uint, num uint, prev *MeasureHeader) (*MeasureHeader, uint8) {
	if prev != nil {
		*off += 1
	}
	mh, flags := s.readMeasureHeader(data, off, num, prev)
	if mh.RepeatClose > -1 {
		mh.RepeatClose -= 1
	}

	if (flags & 0x03) == 0x03 {
		for i := 0; i < 4; i++ {
			mh.TimeSignature.Beams[i] = readByte(data, off)
		}
	} else {
		mh.TimeSignature.Beams = prev.TimeSignature.Beams
	}

	if (flags & 0x10) == 0 {
		*off += 1
	}

	mh.TripletFeel = toTripletFeel[int8(readByte(data, off))]
	return mh, flags
}

func (s *Song) readRepeatAlternative(data []byte, off *uint) uint8 {
	v := uint16(readByte(data, off))
	var existingAlternative uint16
	for i := len(s.MeasureHeaders); i > 0; i-- {
		if s.MeasureHeaders[i].RepeatOpen {
			break
		}
		existingAlternative |= uint16(s.MeasureHeaders[i].RepeatAlternative)
	}
	return uint8(((1 << v) - 1) ^ existingAlternative)
}

func (s *Song) readRepeatAlternativeV5(data []byte, off *uint) uint8 {
	return readByte(data, off)
}
