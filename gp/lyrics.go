package gp

type Lyrics struct {
	TrackChoice byte
	Lines       []Line
}

type Line struct {
	I               byte
	StartingMeasure uint16
	Text            string
}

func (s *Song) readLyrics(data []byte, off *uint) {
	l := Lyrics{
		TrackChoice: byte(readInt(data, off)),
	}
	for i := 0; i < 5; i++ {
		sm := uint16(readInt(data, off))
		l.Lines = append(l.Lines, Line{
			I:               byte(i),
			StartingMeasure: sm,
			Text:            readIntSizeString(data, off),
		})
	}
}
