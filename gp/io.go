package gp

import (
	"encoding/binary"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

func readByte(data []byte, off *uint) byte {
	if uint(len(data)) < *off {
		panic("EOF")
	}
	b := data[*off]
	*off += 1
	return b
}

func readSignedByte(data []byte, off *uint) int8 {
	if uint(len(data)) < *off {
		panic("EOF")
	}
	b := data[*off]
	*off += 1
	return int8(b)
}

func readBool(data []byte, off *uint) bool {
	if uint(len(data)) < *off {
		panic("EOF")
	}
	b := data[*off]
	*off += 1
	return b != 0
}

func readShort(data []byte, off *uint) int16 {
	if uint(len(data)) < *off+2 {
		panic("EOF")
	}
	n := []byte{data[*off], data[*off+1]}
	*off += 2
	return int16(binary.LittleEndian.Uint16(n))
}

func readInt(data []byte, off *uint) int32 {
	if uint(len(data)) < *off+4 {
		panic("EOF")
	}
	n := []byte{data[*off], data[*off+1], data[*off+2], data[*off+3]}
	*off += 4
	return int32(binary.LittleEndian.Uint32(n))
}

func readIntSizeString(data []byte, off *uint) string {
	s := readInt(data, off)
	return readString(data, off, uint(s), 0)
}

func readIntByteSizeString(data []byte, off *uint) string {
	s := readInt(data, off) - 1
	return readByteSizeString(data, off, uint(s))
}

func readByteSizeString(data []byte, off *uint, size uint) string {
	length := uint(readByte(data, off))
	return readString(data, off, size, uint(length))
}

func readString(data []byte, off *uint, size, length uint) string {
	if length == 0 {
		length = size
	}
	dec := charmap.Windows1252.NewDecoder()
	b, err := dec.Bytes(data[*off : *off+length])
	if err != nil {
		s := &strings.Builder{}
		s.Write(data[*off : *off+length])
		*off += size
		return s.String()
	}
	*off += size
	return string(b)
}

var VERSIONS = []Version{
	Version{"FICHIER GUITAR PRO v5.10", [3]uint{5, 1, 0}, false},
	Version{"FICHIER GUITAR PRO v5.10", [3]uint{5, 2, 0}, false},
}

func readVersionString(data []byte, off *uint) Version {
	v := Version{
		Data: readByteSizeString(data, off, 30),
	}

	for _, x := range VERSIONS {
		if v.Data == x.Data {
			v.Number = x.Number
			v.Clipboard = x.Clipboard
			break
		}
	}
	return v
}
