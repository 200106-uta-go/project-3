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
	_, err := os.Stat(PROFILES + profileName + ".yaml")
	if os.IsNotExist(err) {
		fmt.Printf("Profile %s.yaml not found in %s\n", profileName, PROFILES)
		return
	}
	for {
		fmt.Printf("Profile %s.yaml will be removed.\nAre you sure you want to continue (Y/n)? ", profileName)
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
			err = cmd.Run()
			if err != nil {
				fmt.Printf("Failed to remove %s\n", profileName)
			} else {
				fmt.Printf("Profile %s has been removed\n", profileName)
			}
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
		fmt.Printf("All profiles will be removed.\nAre you sure you want to continue (Y/n)?")
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
			err = cmd.Run()
			if err != nil {
				fmt.Println("Failed to remove all profiles.")
			} else {
				fmt.Println("All profiles have been removed.")
			}
			return
		} else if answer == "n" || answer == "N" {
			return
		} else {
			fmt.Println("Invalid response")
		}
	}
}
