package gp

import "strings"

type Point struct{ X, Y uint16 }

type Padding struct{ Left, Right, Top, Bottom uint16 }

type PageSetup struct {
	PageSize            Point
	PageMargin          Padding
	ScoreSizeProportion float32
	HeaderAndFooter     uint16
	Title               string
	Subtitle            string
	Artist              string
	Album               string
	Words               string
	Music               string
	WordsAndMusic       string
	Copyright           string
	PageNumber          string
}

func (s *Song) readPageSetup(data []byte, off *uint) {
	ps := PageSetup{}
	ps.PageSize.X = uint16(readInt(data, off))
	ps.PageSize.Y = uint16(readInt(data, off))
	ps.PageMargin.Left = uint16(readInt(data, off))
	ps.PageMargin.Right = uint16(readInt(data, off))
	ps.PageMargin.Top = uint16(readInt(data, off))
	ps.PageMargin.Bottom = uint16(readInt(data, off))
	ps.ScoreSizeProportion = float32(readInt(data, off)) / 100.0
	ps.HeaderAndFooter = uint16(readShort(data, off))
	ps.Title = readIntSizeString(data, off)
	ps.Subtitle = readIntSizeString(data, off)
	ps.Artist = readIntSizeString(data, off)
	ps.Album = readIntSizeString(data, off)
	ps.Words = readIntSizeString(data, off)
	ps.Music = readIntSizeString(data, off)
	ps.WordsAndMusic = readIntSizeString(data, off)
	c := &strings.Builder{}
	c.WriteString(readIntSizeString(data, off))
	c.WriteString("\n")
	c.WriteString(readIntSizeString(data, off))
	ps.Copyright = c.String()
	ps.PageNumber = readIntSizeString(data, off)
	s.PageSetup = ps
}
