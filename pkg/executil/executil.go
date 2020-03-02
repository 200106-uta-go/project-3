package executil

import (
	"fmt"
	"os/exec"
	"strings"
)

// // ExecHandler handles script commands using sh
// func ExecHandler(bashInput string, errTryCount, errWaitCount int) (int, error) {
// 	// breaks input into lines
// 	s := strings.Split(bashInput, "\n")

// 	// iterate through lines
// 	for i, ln := range s {
// 		// execute current line
// 		er := execLine(ln)

// 		// if line error occurs, retry it errTryCount times
// 		if er != nil {
// 			for i := 0; i < errTryCount; i++ {
// 				// before retry, sleep for errWaitCount seconds
// 				time.Sleep(time.Second * time.Duration(errWaitCount))
// 				// attempt to execute again
// 				er = execLine(ln)
// 				// success will break out of the error catching loop
// 				if er == nil {
// 					break
// 				}
// 			}
// 			// if an error occurred more than errTryCount times, return i, er
// 			return i, er
// 		}
// 	}
// 	// all good
// 	return -1, nil
// }

// ExecHandler handles script commands using sh
func ExecHandler(bashInput string) {
	// breaks input into lines
	s := strings.Split(bashInput, "\n")

	for _, line := range s {
		line = strings.TrimSpace(line)
		execLine(line)
	}
}

// func execLine(bashLn string) error {
// 	split := strings.Split(bashLn, " ")
// 	out, err := exec.Command(split[0], split[1:]...).Output()
// 	fmt.Println(out)
// 	return err
// }

func execLine(line string) {
	split := strings.Split(line, " ")
	cmd := exec.Command(split[0], split[1:]...)
	count := 0

	//loops the command a max of 3 times if command fails
	for count < 3 {

		//run command and output to console
		byte, err := cmd.Output()
		fmt.Println(string(byte))

		//exit for loop if no error
		if err == nil {
			break
		}
		if err != nil {
			panic(err)
		}
		count++
	}
}
