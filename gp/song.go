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

	Tracks         []Track
	MeasureHeaders []MeasureHeader
	Channels       []MidiChannel
	Lyrics         Lyrics
	Tempo          int16
	HideTempo      bool
	TempoName      string
	Key            KeySignature

	TripletFeel  TripletFeel
	MasterEffect RseMasterEffect

	PageSetup PageSetup
}
