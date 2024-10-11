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

	fmt.Printf("Switched to Git profile: %s\n", name)
	fmt.Printf("Username: %s\n", profile.Username)
	fmt.Printf("Email: %s\n", profile.Email)
}

func setGitConfig(key, value string) {
	cmd := exec.Command("git", "config", "--global", key, value)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error setting %s: %v\n", key, err)
		os.Exit(1)
	}
}