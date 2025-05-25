package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch [profile name]",
	Short: "Switch to a specific Git profile",
	Args:  cobra.ExactArgs(1),
	Run:   switchProfile,
}

func switchProfile(cmd *cobra.Command, args []string) {
	name := args[0]
	profile, exists := profiles[name]
	if !exists {
		fmt.Printf("Profile '%s' does not exist.\n", name)
		return
	}

	setGitConfig("user.name", profile.Username)
	setGitConfig("user.email", profile.Email)

	// configure per-profile SSH key via core.sshCommand
	if profile.SSHKey != "" {
		// Git 2.10+ supports this
		sshCmd := fmt.Sprintf("ssh -i %s", profile.SSHKey)
		setGitConfig("core.sshCommand", sshCmd)
	} else {
		// fall back to default SSH key: remove custom setting if previously set
		exec.Command("git", "config", "--global", "--unset", "core.sshCommand").Run()
	}

	fmt.Printf("Switched to Git profile: %s\n", name)
	fmt.Printf("Username: %s\n", profile.Username)
	fmt.Printf("Email: %s\n", profile.Email)
	fmt.Printf("SSH key: %s\n", profile.SSHKey)
}

func setGitConfig(key, value string) {
	cmd := exec.Command("git", "config", "--global", key, value)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error setting %s: %v\n", key, err)
		os.Exit(1)
	}
}
