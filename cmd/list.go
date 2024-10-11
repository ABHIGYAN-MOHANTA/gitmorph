package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all Git profiles",
	Run:   listProfiles,
}

func listProfiles(cmd *cobra.Command, args []string) {
	fmt.Println("Available Git profiles:")
	for name, profile := range profiles {
		fmt.Printf("- %s (Username: %s, Email: %s)\n", name, profile.Username, profile.Email)
	}
}