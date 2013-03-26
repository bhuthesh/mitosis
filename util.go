// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package mitosis

import (
	"fmt"
	"os"
	"time"
)

// LogFile determines the location to which log entries are written.
// Set this to nil, to turn logging off.
var LogFile = os.Stdout

// Log outputs log information to a previously specified target file.
// It only outputs data if the Verbose flag is set to true.
func Log(f string, argv ...interface{}) {
	if LogFile == nil {
		return
	}

	stamp := time.Now().UTC().Format("2006-01-02 15:04:05")
	fmt.Fprintf(LogFile, "%s %s\n", stamp, fmt.Sprintf(f, argv...))
}

// check panics with the given error if it is not nil.
func check(err error) {
	if err != nil {
		panic(err)
	}
}
