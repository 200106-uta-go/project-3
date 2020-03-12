package kreate

// ListProfiles displays all .yaml files in the profile directory.
func ListProfiles() string {
	str, _ := shellCommand("ls -a | grep .yaml", PROFILES)
	return str
}