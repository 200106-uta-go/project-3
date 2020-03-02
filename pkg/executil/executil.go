package executil

import (
	"os/exec"
	"strings"
	"time"
)

// ExecHandler handles script commands using sh
func ExecHandler(bashInput string, errTryCount, errWaitCount int) (int, error) {
	// breaks input into lines
	s := strings.Split(bashInput, "\n")

	// iterate through lines
	for i, ln := range s {
		// execute current line
		er := execLine(ln)

		// if line error occurs, retry it errTryCount times
		if er != nil {
			for i := 0; i < errTryCount; i++ {
				// before retry, sleep for errWaitCount seconds
				time.Sleep(time.Second * time.Duration(errWaitCount))
				// attempt to execute again
				er = execLine(ln)
				// success will break out of the error catching loop
				if er == nil {
					break
				}
			}
			// if an error occurred more than errTryCount times, return i, er
			return i, er
		}
	}
	// all good
	return -1, nil
}

func execLine(bashLn string) error {
	_, err := exec.Command("/bin/sh", "-c", "\""+bashLn+"\"").Output()
	return err
}
