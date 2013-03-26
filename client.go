// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package mitosis

import (
	"fmt"
	"net"
)

// spawnClient spawns a new client connection to the given server port.
func spawnClient(port uint, sf StateFunc) (err error) {
	defer func() {
		if x := recover(); x != nil {
			if e, ok := x.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("%v", x)
			}
		}
	}()

	conn, err := net.Dial(protocol, fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	defer conn.Close()

	writeRaw(conn, magick)

	var state State
	state.read(conn)
	sf(&state)

	return nil
}
