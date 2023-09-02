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

func (s *Song) readLyrics(data []byte, seek *uint) {
	l := Lyrics{
		TrackChoice: byte(readInt(data, seek)),
	}
	for i := uint8(0); i < 5; i++ {
		sm := uint16(readInt(data, seek))
		l.Lines = append(l.Lines, Line{
			I:               i,
			StartingMeasure: sm,
			Text:            readIntSizeString(data, seek),
		})
	}
	s.Lyrics = l
}
