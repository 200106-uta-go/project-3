package kreate

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// RemoveProfile removes a specified profile from the directory.
func RemoveProfile(profileName string) {
	//check if profileName has an extension, if not add .yaml
	if !strings.HasSuffix(profileName, ".yaml") && !strings.HasSuffix(profileName, ".yml") {
		profileName += ".yaml"
	}

	_, err := os.Stat(PROFILES + profileName)
	if os.IsNotExist(err) {
		fmt.Printf("Profile %s not found in %s\n", profileName, PROFILES)
		return
	}
	for {
		fmt.Printf("Profile %s will be removed.\nAre you sure you want to continue (Y/n)? ", profileName)
		reader := bufio.NewReader(os.Stdin)
		answer, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		answer = strings.Replace(answer, "\n", "", -1)
		if answer == "y" || answer == "Y" {
			cmd := exec.Command("rm", PROFILES+profileName)
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
