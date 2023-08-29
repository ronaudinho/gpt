package gp

import (
	"fmt"

	"golang.org/x/mod/semver"
)

func (s *Song) ReadGP5(data []byte) {
	off := new(uint)
	s.Version = readVersionString(data, off)
	s.readClipboard(data, off) // TODO can skip since false now
	s.readInfo(data, off)
	s.readLyrics(data, off)
	s.readRseMasterEffect(data, off)
	s.readPageSetup(data, off)
	s.TempoName = readIntSizeString(data, off)
	s.Tempo = int16(readInt(data, off))
	sv := fmt.Sprintf("v%d.%d.%d",
		s.Version.Number[0],
		s.Version.Number[1],
		s.Version.Number[2],
	)
	if semver.Compare(sv, "v5.0.0") > 0 {
		s.HideTempo = readBool(data, off)
	}
}

func (s *Song) readInfo(data []byte, off *uint) {
	s.Name = readIntByteSizeString(data, off)
	s.Subtitle = readIntByteSizeString(data, off)
	s.Artist = readIntByteSizeString(data, off)
	s.Album = readIntByteSizeString(data, off)
	s.Words = readIntByteSizeString(data, off)
	if s.Version.Number[0] < 5 {
		s.Author = s.Words
	} else {
		s.Author = readIntByteSizeString(data, off)
	}
	s.Copyright = readIntByteSizeString(data, off)
	s.Writer = readIntByteSizeString(data, off)
	s.Instructions = readIntByteSizeString(data, off)
	nc := readInt(data, off)
	if nc > 0 {
		for i := int32(0); i < nc; i++ {
			// TODO ordering maybe
			s.Notice = append(s.Notice, readIntByteSizeString(data, off))
		}
	}
}
