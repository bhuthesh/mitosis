// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package mitosis

import (
	"encoding/binary"
	"io"
)

// check panics with the given error if it is not nil.
func check(err error) {
	if err != nil {
		panic(err)
	}
}

var endian = binary.LittleEndian

func readU32(r io.Reader) uint32 {
	var v uint32
	check(binary.Read(r, endian, &v))
	return v
}

func readU64(r io.Reader) uint64 {
	var v uint64
	check(binary.Read(r, endian, &v))
	return v
}

func readRaw(r io.Reader, size uint32) []byte {
	if size == 0 {
		return nil
	}

	data := make([]byte, size)
	_, err := r.Read(data)
	check(err)
	return data
}

func readByteSlice(r io.Reader) []byte {
	size := readU32(r)
	return readRaw(r, size)
}

func readStringSlice(r io.Reader) []string {
	size := readU32(r)
	list := make([]string, size)

	for i := range list {
		list[i] = string(readByteSlice(r))
	}

	return list
}
