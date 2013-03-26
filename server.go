// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package mitosis

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"strconv"
)

const (
	defaultPort = 32000 // First port at which we start looking for available sockets.
	protocol    = "tcp" // Transport protocol being used.
)

var magick = []byte("MTSS") // Protocol header, identifying a valid mitosis session.

// spawnServer opens the sever's listening socket.
func spawnServer(state *State, done chan<- bool) (uint, error) {
	var listener net.Listener
	var port uint
	var err error

	port = defaultPort

	// open a port and get the port number
	listener, err = net.Listen(protocol, ":0")
	if err != nil {
		return 0, errors.New("Failed to open a listen port")
	}

	laddr := listener.Addr().String()
	_, portstr, err := net.SplitHostPort(laddr)
	if err != nil {
		return 0, fmt.Errorf("Failed to get port from listen port address string: %s", laddr)
	}

	n, err := strconv.ParseUint(portstr, 10, 16)
	if err != nil {
		return 0, fmt.Errorf("Bad port number string: %s", portstr)
	}

	port = uint(n)

	go func() {
		defer listener.Close()

		for {
			client, err := listener.Accept()
			if err != nil {
				return
			}

			handleClient(client, state, done)
		}
	}()

	return port, nil
}

// handleClient handles an incoming client connection.
func handleClient(client net.Conn, state *State, done chan<- bool) {
	defer func() {
		client.Close()
		recover()
	}()

	// Make sure the protocol header is there and valid.
	hdr := readRaw(client, uint32(len(magick)))

	if !bytes.Equal(hdr, magick) {
		return
	}

	// Send pending state information.
	state.write(client)

	// Signal completion.
	done <- true
}
