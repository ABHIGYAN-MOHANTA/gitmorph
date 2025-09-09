package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var activateCmd = &cobra.Command{
	Use:   "activate [profile name]",
	Short: "Activate a Git profile for the current project (writes .gitmorph file and switches)",
	Args:  cobra.ExactArgs(1),
	Run:   activateProfile,
}

func init() {
	RootCmd.AddCommand(activateCmd)
}

func activateProfile(cmd *cobra.Command, args []string) {
	name := args[0]

	// Verify profile exists
	profile, exists := profiles[name]
	if !exists {
		fmt.Printf("Profile '%s' does not exist.\n", name)
		os.Exit(1)
	}

	// Write .gitmorph file in current directory
	data := map[string]string{"profile": name}
	bytes, _ := json.MarshalIndent(data, "", "  ")
	path := filepath.Join(".", ".gitmorph")
	if err := os.WriteFile(path, bytes, 0644); err != nil {
		fmt.Printf("Error writing .gitmorph file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Activated profile '%s' for this project (written to .gitmorph).\n", name)

	// Immediately switch Git config globally
	if err := switchToProfile(profile.Name, false); err != nil {
		fmt.Printf("Error switching to profile '%s': %v\n", profile.Name, err)
		os.Exit(1)
	}
}
