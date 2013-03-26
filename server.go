// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package mitosis

import (
	"bytes"
	"errors"
	"fmt"
	"net"
)

const (
	defaultPort = 32000 // First port at which we start looking for available sockets.
	protocol    = "tcp" // Transport protocol being used.
)

var magick = []byte("MTSS") // Protocol header, identifying a valid mitosis session.

// spawnServer opens the sever's listening socket.
func spawnServer(state *State, done chan<- bool) (uint16, error) {
	var listener net.Listener
	var port uint16
	var err error

	port = defaultPort

	// Find an open port, the brute-force way.
	for {
		listener, err = net.Listen(protocol, fmt.Sprintf(":%d", port))

		if err == nil {
			break
		}

		port++

		// Wrapped around uint16 capacity?
		if port == 0 {
			return 0, errors.New("No valid port number could be found.")
		}
	}

	go func() {
		defer listener.Close()

		Log("[mitosis] Listening on :%d...", port)

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

		if err := recover(); err != nil {
			Log("%v", err)
		}
	}()

	Log("[mitosis] Client %s connected", client.RemoteAddr())

	Log("[mitosis] Receiving protocol header...")
	// Make sure the protocol header is there and valid.
	hdr := readRaw(client, uint32(len(magick)))

	if !bytes.Equal(hdr, magick) {
		return
	}

	Log("[mitosis] Sending application state...")
	// Send pending state information.
	state.Write(client)

	// Signal completion.
	Log("[mitosis] Done.")
	done <- true
}
