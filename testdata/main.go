// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

/*
This program shows how a client can use the mitosis service to issue
automatic relaunches with preservation of application state across sessions.

WARNING: This code is for demonstration purposes only.
Running it, will cause it to spawn a new instance of itself every 5 seconds.
This can only be stopped by issueing a `pkill testdata` command from a shell.
*/
package main

import (
	"flag"
	"fmt"
	"github.com/jteeuwen/mitosis"
	"os"
	"time"
)

func main() {
	flag.Parse()

	// Redirect log output to a file.
	mitosis.LogFile, _ = os.Create(fmt.Sprintf("%d.log", time.Now().Unix()))
	defer mitosis.LogFile.Close()

	// Initialize mitosis and give it a StateFunc handler. This will be called
	// with state data from a previous session if applicable.
	// Note that this must be done /after/ the mandatory flag.Parse() call.
	err := mitosis.Init(onState)
	if err != nil {
		mitosis.Log("Init error: %v", err)
		return
	}

	// Wait a bit; pretend we are doing important program things.
	<-time.After(5e9)

	// Create our application state structure.
	state := &mitosis.State{
		Commandline: nil,                    // Command line arguments we need to pass on.
		Files:       nil,                    // File handles which need to be inherited.
		Data:        []byte("Sample data."), // Custom data, marshalled to a byte slice.
	}

	// Split ourselves off into a new session.
	done, err := mitosis.Split(state)
	if err != nil {
		mitosis.Log("Split error: %v", err)
		return
	}

	// Sit and wait for us to receive the OK to shut down.
	<-done
}

// onState handles application state sent by an old program session.
func onState(state *mitosis.State) {
	mitosis.Log("onState: %v", state.Data)
}
