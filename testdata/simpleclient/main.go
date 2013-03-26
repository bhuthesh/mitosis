// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

/*
This program shows how a client can use the mitosis service to issue
automatic relaunches with preservation of application state across sessions.

WARNING: This code is for demonstration purposes only.
Running it, will cause it to spawn a new instance of itself every 5 seconds.
This can only be stopped by issueing a `pkill simpleclient` command from a shell.
*/
package main

import (
	"flag"
	"github.com/jteeuwen/mitosis"
	"log"
	"os"
	"time"
)

var (
	logger *log.Logger
)

func main() {
	flag.Parse()

	logfile := initLog()
	defer logfile.Close()

	// Initialize mitosis and give it a StateFunc handler. This will be called
	// with state data from a previous session if applicable.
	// Note that this must be done /after/ the mandatory flag.Parse() call.
	_, err := mitosis.Init(onState)
	if err != nil {
		logger.Fatalf("Init error: %v", err)
	}

	// Wait a bit; pretend we are doing important program things.
	<-time.After(5e9)

	// Split ourselves off into a new session.
	state := &mitosis.State{
		Data: []byte("Sample data."),
	}
	done, err := mitosis.Split(nil, state)
	if err != nil {
		logger.Fatalf("Split error: %v", err)
	}

	// Sit and wait for us to receive the OK to shut down.
	<-done
}

// onState handles application state sent by an old program session.
func onState(state *mitosis.State) {
	logger.Printf("onState: %s", state.Data)
}

// initLog initialize a logger to write to an external file.
func initLog() *os.File {
	logfile, err := os.OpenFile("debug.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}

	logger = log.New(logfile, "", log.LstdFlags)
	return logfile
}
