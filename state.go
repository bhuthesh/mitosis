// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package mitosis

import (
	"io"
	"os"
)

type StateFunc func(*State)

// State defines the applicaiton state which should be transfered
// to a new program session.
type State struct {
	Commandline []string   // Optional commandline arguments for the new session.
	Data        []byte     // Custom blob of state information.
	Files       []*os.File // List of file descriptors that need to be inherited.
}

func (sd *State) Write(w io.Writer) {
	writeStringSlice(w, sd.Commandline)
	writeByteSlice(w, sd.Data)
	writeU32(w, uint32(len(sd.Files)))

	for _, f := range sd.Files {
		writeUintptr(w, f.Fd())
	}
}

func (sd *State) Read(r io.Reader) {
	sd.Commandline = readStringSlice(r)
	sd.Data = readByteSlice(r)
	sd.Files = make([]*os.File, readU32(r))

	for i := range sd.Files {
		fd := readUintptr(r)
		sd.Files[i] = os.NewFile(fd, "")
	}
}
