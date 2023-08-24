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
	Notice       []string
	Lyrics       Lyrics
	MasterEffect RseMasterEffect

	Date        string
	Transcriber string
	Comments    string
}
