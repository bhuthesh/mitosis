// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package mitosis

import (
	"encoding/binary"
	"io"
)

func writeString(w io.Writer, v string) {}

func writeU32(w io.Writer, v uint32) {
	check(binary.Write(w, endian, v))
}

func writeU64(w io.Writer, v uint64) {
	check(binary.Write(w, endian, v))
}

func writeRaw(w io.Writer, v []byte) {
	_, err := w.Write(v)
	check(err)
}

func writeByteSlice(w io.Writer, v []byte) {
	writeU32(w, uint32(len(v)))
	writeRaw(w, v)
}

func writeStringSlice(w io.Writer, v []string) {
	writeU32(w, uint32(len(v)))

	for i := range v {
		writeByteSlice(w, []byte(v[i]))
	}
}
