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

func createNewProfile(cmd *cobra.Command, args []string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter profile name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

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

	profiles[name] = Profile{
		Name:     name,
		Username: username,
		Email:    email,
		SSHKey:   sshKey,
	}

	saveProfiles()
	fmt.Printf("Profile '%s' created successfully.\n", name)
}
