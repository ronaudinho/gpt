package gp

type Song struct {
	Version      Version
	Name         string
	Subtitle     string
	Artist       string
	Album        string
	Words        string
	Author       string
	Copyright    string
	Writer       string
	Instructions string
	Date         string
	Transcriber  string
	Comments     string
	Notice       []string

	Lyrics    Lyrics
	Tempo     int16
	TempoName string
	HideTempo bool
	Key       KeySignature
	Channels  []MidiChannel

	MasterEffect RseMasterEffect

	PageSetup PageSetup
}
