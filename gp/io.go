package gp

import (
	"encoding/binary"

	"golang.org/x/text/encoding/charmap"
)

func readByte(data []byte, seek *uint) byte {
	if uint(len(data)) < *seek {
		panic("EOF")
	}
	b := data[*seek]
	*seek += 1
	return b
}

func readSignedByte(data []byte, seek *uint) byte {
	return 0
}

func readBool(data []byte, seek *uint) bool {
	return false
}

func readShort(data []byte, seek *uint) uint16 {
	return 0
}

func readInt(data []byte, seek *uint) uint32 {
	if uint(len(data)) < *seek+4 {
		panic("EOF")
	}
	n := []byte{data[*seek], data[*seek+1], data[*seek+2], data[*seek+3]}
	*seek += 4
	return binary.LittleEndian.Uint32(n)
}

func readIntByteSizeString(data []byte, seek *uint) string {
	s := readInt(data, seek) - 1
	return readByteSizeString(data, seek, uint(s))
}

func readByteSizeString(data []byte, seek *uint, size uint) string {
	length := uint(readByte(data, seek))
	return readString(data, seek, size, uint(length))
}

func readString(data []byte, seek *uint, size, length uint) string {
	if length == 0 {
		length = size
	}
	dec := charmap.Windows1252.NewDecoder()
	b, err := dec.Bytes(data[*seek : *seek+length])
	if err != nil {
		panic(err)
	}
	*seek += size
	return string(b)
}

var VERSIONS = []Version{
	Version{"FICHIER GUITAR PRO v5.10", [3]uint{5, 1, 0}, false},
	Version{"FICHIER GUITAR PRO v5.10", [3]uint{5, 2, 0}, false},
}

func readVersionString(data []byte, seek *uint) Version {
	v := Version{
		Data: readByteSizeString(data, seek, 30),
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
