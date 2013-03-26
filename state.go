// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package mitosis

import (
	"io"
	"os"
)

// StateFunc is a callback which is called during library initialization.
// It passes any state information from an old program session into the
// new session.
type StateFunc func(*State)

// State defines the application state that should be transfered
// to a new program session.
type State struct {
	Data  []byte     // Custom blob of state information.
	Files []*os.File // List of file descriptors that need to be inherited.
}

func (sd *State) write(w io.Writer) {
	writeByteSlice(w, sd.Data)
	writeU32(w, uint32(len(sd.Files)))

	for _, f := range sd.Files {
		writeU64(w, uint64(f.Fd()))
	}
}

func (sd *State) read(r io.Reader) {
	sd.Data = readByteSlice(r)
	sd.Files = make([]*os.File, readU32(r))

	for i := range sd.Files {
		fd := uintptr(readU64(r))
		sd.Files[i] = os.NewFile(fd, "")
	}
}
