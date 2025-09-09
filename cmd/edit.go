package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [profile name]",
	Short: "Edit an existing Git profile",
	Args:  cobra.ExactArgs(1),
	Run:   editProfile,
}

func init() {
	RootCmd.AddCommand(editCmd)
}

func editProfile(cmd *cobra.Command, args []string) {
	name := args[0]
	profile, exists := profiles[name]
	if !exists {
		fmt.Printf("Profile '%s' does not exist.\n", name)
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Editing profile '%s'. Leave blank to keep current value.\n", name)

	// Username
	fmt.Printf("Current Username [%s]: ", profile.Username)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		profile.Username = input
	}

	// Email
	fmt.Printf("Current Email    [%s]: ", profile.Email)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		profile.Email = input
	}

	// SSH Key
	fmt.Printf("Current SSH Key  [%s]: ", profile.SSHKey)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input != "" {
		profile.SSHKey = input
	}

	// Default flag
	fmt.Printf("Is this the default profile? (current: %v) [y/N]: ", profile.Default)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))

	if input == "y" || input == "yes" {
		// clear other defaults
		for k, p := range profiles {
			p.Default = false
			profiles[k] = p
		}
		profile.Default = true
	} else if input == "n" || input == "no" {
		// check if this is the only default profile
		isOnlyDefault := true
		for k, p := range profiles {
			if k != name && p.Default {
				isOnlyDefault = false
				break
			}
		}
		if isOnlyDefault {
			fmt.Println("⚠️ Cannot unset default: at least one profile must be default.")
			profile.Default = true
		} else {
			profile.Default = false
		}
	}

	// Save back into map and disk
	profiles[name] = profile
	saveProfiles()

	fmt.Printf("Profile '%s' updated successfully.\n", name)
}
