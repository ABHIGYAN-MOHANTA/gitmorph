package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var defaultCmd = &cobra.Command{
	Use:   "default [profile name]",
	Short: "Set a Git profile as the default",
	Args:  cobra.ExactArgs(1),
	Run:   setDefaultProfile,
}

func init() {
	RootCmd.AddCommand(defaultCmd)
}

func setDefaultProfile(cmd *cobra.Command, args []string) {
	name := args[0]

	// Check if profile exists
	profile, exists := profiles[name]
	if !exists {
		fmt.Printf("Profile '%s' does not exist.\n", name)
		os.Exit(1)
	}

	// Reset all profiles to Default=false
	for key, p := range profiles {
		if p.Default {
			p.Default = false
			profiles[key] = p
		}
	}

	// Set the chosen profile as default
	profile.Default = true
	profiles[name] = profile
	saveProfiles()

	fmt.Printf("Profile '%s' is now the default.\n", name)
}
