package gp

import (
	"fmt"

	"golang.org/x/mod/semver"
)

type RseEqualizer struct {
	Knobs []float32
	Gain  float32
}

type RseMasterEffect struct {
	Volume    float32
	Reverb    float32
	Equalizer RseEqualizer
}

func (s *Song) readRseMasterEffect(data []byte, off *uint) {
	var me RseMasterEffect
	// version >= 5.1.0
	sv := fmt.Sprintf("v%d.%d.%d",
		s.Version.Number[0],
		s.Version.Number[1],
		s.Version.Number[2],
	)
	if semver.Compare(sv, "v5.1.0") < 0 {
		return
	}
	me.Volume = float32(readInt(data, off))
	readInt(data, off) // ???
	me.Equalizer = readRseEqualizer(data, off, 11)
	s.MasterEffect = me
}

func readRseEqualizer(data []byte, off *uint, knobs byte) RseEqualizer {
	var e RseEqualizer
	for i := uint8(0); i < knobs; i++ {
		e.Knobs = append(e.Knobs, unpackVolumeValue(readSignedByte(data, off)))
	}
	return e
}

func unpackVolumeValue(v int8) float32 {
	return float32(-v) / 10.0
}
