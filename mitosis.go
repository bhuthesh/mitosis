// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package mitosis

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

var serverPort uint

func init() {
	flag.UintVar(&serverPort, "mitosis", 0, "For use by mitosis package.")
}

// Init initializes the mitosis library.
// It should be called after flag.Parse() has been called.
//
// If this program session was launched by mitosis, it will launch
// a client connection to the owner process to fetch program state.
//
// It accepts a StateFunc handler, which will be called with the application
// state data from the old program instance.
func Init(sf StateFunc) (bool, error) {
	if serverPort == 0 {
		return false, nil
	}

	return true, spawnClient(serverPort, sf)
}

// Split forks off a new copy of the current application, and hands it the
// given state information. It returns a channel on which mitosis will signal
// if the fork was successful. This would indicate it is safe for the current
// session to shut down.
func Split(commandline []string, state *State) (<-chan bool, error) {
	if state == nil {
		return nil, errors.New("Invalid state")
	}

	done := make(chan bool)
	path := filepath.Clean(os.Args[0])
	path, _ = filepath.Abs(path)

	// Launch the server.
	port, err := spawnServer(state, done)
	if err != nil {
		return nil, err
	}

	// Launch the new program instance.
	argv := append(commandline, "-mitosis", fmt.Sprintf("%d", port))

	cmd := exec.Command(path, argv...)
	cmd.Dir, _ = filepath.Split(path)
	cmd.ExtraFiles = state.Files
	cmd.Env = os.Environ()

	return done, cmd.Start()
}
