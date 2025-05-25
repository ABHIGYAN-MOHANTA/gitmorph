package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [profile name]",
	Short: "Delete a Git profile",
	Args:  cobra.ExactArgs(1),
	Run:   deleteProfile,
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}

func deleteProfile(cmd *cobra.Command, args []string) {
	name := args[0]
	// check existence
	if _, exists := profiles[name]; !exists {
		fmt.Printf("Profile '%s' does not exist.\n", name)
		os.Exit(1)
	}

	// remove the profile and save
	delete(profiles, name)
	saveProfiles()

	fmt.Printf("Profile '%s' deleted successfully.\n", name)
}
