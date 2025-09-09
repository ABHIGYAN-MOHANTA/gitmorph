package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all Git profiles",
	Run:   listProfiles,
}

func init() {
	RootCmd.AddCommand(listCmd)
}

func listProfiles(cmd *cobra.Command, args []string) {
	var activeProfile string

	// Check for project-specific .gitmorph
	if data, err := os.ReadFile(".gitmorph"); err == nil {
		var f map[string]string
		if jsonErr := json.Unmarshal(data, &f); jsonErr == nil {
			activeProfile = f["profile"]
		}
	}

	fmt.Println("Available Git profiles:")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "NAME\tUSERNAME\tEMAIL\tSSH KEY\tFLAGS")

	for name, profile := range profiles {
		flags := ""
		if profile.Default {
			flags += "[default] "
		}
		if activeProfile != "" && name == activeProfile {
			flags += "[active] "
		}

		fmt.Fprintf(
			w,
			"%s\t%s\t%s\t%s\t%s\n",
			name, profile.Username, profile.Email, profile.SSHKey, flags,
		)
	}

	w.Flush()
}
