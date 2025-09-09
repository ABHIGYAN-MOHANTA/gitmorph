package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new Git profile",
	Run:   createNewProfile,
}

func init() {
	RootCmd.AddCommand(newCmd)
}

func createNewProfile(cmd *cobra.Command, args []string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter profile name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	if _, exists := profiles[name]; exists {
		fmt.Printf("Profile '%s' already exists. Use 'edit %s' to modify it.\n", name, name)
		return
	}

	fmt.Print("Enter Git username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("Enter Git email: ")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)

	fmt.Print("Enter SSH private key path (leave blank for default ~/.ssh/id_ed25519): ")
	sshKey, _ := reader.ReadString('\n')
	sshKey = strings.TrimSpace(sshKey)
	if sshKey == "" {
		sshKey = "~/.ssh/id_ed25519"
	}

	// Default to false
	newProfile := Profile{
		Name:     name,
		Username: username,
		Email:    email,
		SSHKey:   sshKey,
		Default:  false,
	}

	// Ask if this should be default
	fmt.Print("Set this profile as default? (y/N): ")
	resp, _ := reader.ReadString('\n')
	resp = strings.TrimSpace(strings.ToLower(resp))

	if resp == "y" || resp == "yes" {
		// Clear any old default
		for k, p := range profiles {
			p.Default = false
			profiles[k] = p
		}
		newProfile.Default = true
	}

	profiles[name] = newProfile
	saveProfiles()

	fmt.Printf("Profile '%s' created successfully", name)
	if newProfile.Default {
		fmt.Print(" and set as default.\n")
	} else {
		fmt.Print(".\n")
	}
}
