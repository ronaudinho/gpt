package gp

func (s *Song) ReadGP5(data []byte) {
	off := new(uint)
	s.Version = readVersionString(data, off)
	s.readClipboard(data, off) // TODO can skip since false now
	s.readInfo(data, off)
	s.readLyrics(data, off)
	s.readRseMasterEffect(data, off)
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
		for i := uint32(0); i < nc; i++ {
			// TODO ordering maybe
			s.Notice = append(s.Notice, readIntByteSizeString(data, off))
		}
	}
}
