package executil

import (
	"os/exec"
	"strings"
)

// ExecHandler handles script commands using sh
func ExecHandler(bashIn string, index int) (int, error) {
	s := strings.Split(bashIn, "\n")
	lens := len(s)
	for i := index; i < lens; i++ {
		_, err := exec.Command("/bin/sh", "-c", "\""+s[i]+"\"").Output()
		if err != nil {
			return i, err
		}
	}
	return lens, nil
}
