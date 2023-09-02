package gp

const DEFAULT_PERCUSSION_CHANNEL uint8 = 9

type MidiChannel struct {
	Channel       uint8
	EffectChannel uint8
	Bank          uint8
	Volume        int8
	Balance       int8
	Chorus        int8
	Reverb        int8
	Phaser        int8
	Tremolo       int8
	Instrument    int32
}

var DefaultMidiChannel = MidiChannel{
	EffectChannel: 1,
	Volume:        104,
	Balance:       64,
	Instrument:    25,
}

func (mc *MidiChannel) isPercussionChannel(ins int32) bool {
	return (mc.Channel % 16) == DEFAULT_PERCUSSION_CHANNEL
}

func (mc *MidiChannel) setInstrument(ins int32) int32 {
	if mc.Instrument == -1 && mc.isPercussionChannel(ins) {
		return 0
	}
	return ins
}

func (s *Song) readMidiChannels(data []byte, off *uint) {
	for i := uint8(0); i < 64; i++ {
		s.Channels = append(s.Channels, s.readMidiChannel(data, off, i))
	}
}

func (s *Song) readMidiChannel(data []byte, off *uint, ch uint8) MidiChannel {
	instrument := readInt(data, off)
	mc := DefaultMidiChannel
	mc.Channel = ch
	mc.EffectChannel = 1
	mc.Volume = readSignedByte(data, off)
	mc.Balance = readSignedByte(data, off)
	mc.Chorus = readSignedByte(data, off)
	mc.Reverb = readSignedByte(data, off)
	mc.Phaser = readSignedByte(data, off)
	mc.Tremolo = readSignedByte(data, off)
	mc.Instrument = mc.setInstrument(instrument)
	*off += 2
	return mc
}
