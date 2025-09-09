package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var deactivateCmd = &cobra.Command{
	Use:   "deactivate",
	Short: "Remove the active Git profile for this project (delete .gitmorph file and switch to default)",
	Run:   deactivateProfile,
}

func init() {
	RootCmd.AddCommand(deactivateCmd)
}

func deactivateProfile(cmd *cobra.Command, args []string) {
	path := filepath.Join(".", ".gitmorph")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("No .gitmorph file found in this project. Nothing to deactivate.")
		return
	}

	if err := os.Remove(path); err != nil {
		fmt.Printf("Error removing .gitmorph: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Project-specific profile deactivated. Falling back to default profile.")

	// Switch to default profile globally
	defaultFound := false
	for name, profile := range profiles {
		if profile.Default {
			if err := switchToProfile(name, false); err != nil {
				fmt.Printf("Error switching to default profile '%s': %v\n", name, err)
				os.Exit(1)
			}
			defaultFound = true
			break
		}
	}

	if !defaultFound {
		fmt.Println("⚠️ No default profile set. Git config unchanged.")
	}
}
