package executil

import (
	"os"
	"os/exec"
	"strings"
)

// ExecHandler handles script commands using sh
func ExecHandler(bashIn string) ([]byte, error) {
	out, err := exec.Command("/bin/sh", "-c", "\"" + bashIn + "\"").Output()
	return out, err
}