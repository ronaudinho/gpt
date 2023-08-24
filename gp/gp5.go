package gp

func (s *Song) ReadGP5(data []byte) {
	seek := new(uint)
	s.Version = readVersionString(data, seek)
	s.readClipboard(data, seek) // TODO can skip since false now
	s.readInfo(data, seek)
}

func (s *Song) readInfo(data []byte, seek *uint) {
	s.Name = readIntByteSizeString(data, seek)
	s.Subtitle = readIntByteSizeString(data, seek)
	s.Album = readIntByteSizeString(data, seek)
	s.Words = readIntByteSizeString(data, seek)
	if s.Version.Number[0] < 5 {
		s.Author = s.Words
	} else {
		s.Author = readIntByteSizeString(data, seek)
	}
	s.Copyright = readIntByteSizeString(data, seek)
	s.Writer = readIntByteSizeString(data, seek)
	s.Instructions = readIntByteSizeString(data, seek)
}
