package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

type Profile struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var profiles map[string]Profile
var configFile string

var RootCmd = &cobra.Command{
	Use:   "gitmorph",
	Short: "A CLI tool to seamlessly switch between Git identities",
	Long:  `GitMorph is a command-line tool that helps you manage and effortlessly switch between different Git accounts on your local machine.`,
}

func init() {
	homeDir, _ := os.UserHomeDir()
	configFile = filepath.Join(homeDir, ".gitmorph.json")

	RootCmd.AddCommand(newCmd)
	RootCmd.AddCommand(switchCmd)
	RootCmd.AddCommand(listCmd)

	RootCmd.CompletionOptions.DisableDefaultCmd = true

	loadProfiles()
}

func loadProfiles() {
	profiles = make(map[string]Profile)

	data, err := ioutil.ReadFile(configFile)
	if err == nil {
		json.Unmarshal(data, &profiles)
	}
}

func saveProfiles() {
	data, _ := json.MarshalIndent(profiles, "", "  ")
	ioutil.WriteFile(configFile, data, 0644)
}