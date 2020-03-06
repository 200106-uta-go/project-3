package kreate

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

/*
## kreate remove <profile name> (--all, -a)
1. The specified profile must be deleted (confirmation up to implementer).
2. (up to implementer) The --all, -a flag must delete all flag.
*/

// RemoveProfile removes a specified profile from the directory.
func RemoveProfile(profileName string) {
	for {
		fmt.Printf("%s profile will be removed.\nAre you sure you want to continue (Y/n)? ", profileName)
		reader := bufio.NewReader(os.Stdin)
		answer, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		answer = strings.Replace(answer, "\n", "", -1)
		if answer == "y" || answer == "Y" {
			cmd := exec.Command("rm", PROFILES+profileName+".yaml")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()

			fmt.Printf("%s profile has been removed\n", profileName)
			return
		} else if answer == "n" || answer == "N" {
			return
		} else {
			fmt.Println("Invalid response")
		}
	}
}

// RemoveAllProfiles removes all profiles from the directory.
func RemoveAllProfiles() {
	for {
		fmt.Printf("All profiles will be removed.\nAre you sure you want to continue (Y/n)? ")
		reader := bufio.NewReader(os.Stdin)
		answer, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		answer = strings.Replace(answer, "\n", "", -1)
		if answer == "y" || answer == "Y" {
			cmd := exec.Command("rm", PROFILES+"*.yaml")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()

			fmt.Println("All profiles have been removed")
			return
		} else if answer == "n" || answer == "N" {
			return
		} else {
			fmt.Println("Invalid response")
		}
	}
}
