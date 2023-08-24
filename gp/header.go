package gp

type Version struct {
	Data      string
	Number    [3]uint
	Clipboard bool
}

type Clipboard struct {
	// TODO
}

func (s *Song) readClipboard(data []byte, seek *uint) Clipboard {
	var c Clipboard
	if !s.Version.Clipboard {
		return c
	}
	return c
}
