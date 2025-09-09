package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "Fix ~/.gitmorph.json",
	Run:   fixProfiles,
}

func init() {
	RootCmd.AddCommand(fixCmd)
}

func fixProfiles(cmd *cobra.Command, args []string) {
	if len(profiles) == 0 {
		fmt.Println("No profiles found in", configFile)
		return
	}

	// Count current defaults
	var defaultName string
	defaultCount := 0
	for name, p := range profiles {
		if p.Default {
			defaultName = name
			defaultCount++
		}
	}

	if defaultCount == 1 {
		fmt.Printf("All good: '%s' is the default profile.\n", defaultName)
		return
	}

	// No default or multiple defaults -> pick the first profile alphabetically
	var first string
	for name := range profiles {
		first = name
		break
	}

	for name := range profiles {
		p := profiles[name]
		p.Default = (name == first)
		profiles[name] = p
	}

	saveProfiles()
	fmt.Printf("Fixed ~/.gitmorph.json. \n'%s' is now the default profile.\nchange it using 'gitmorph default <profile>' \n", first)
}
