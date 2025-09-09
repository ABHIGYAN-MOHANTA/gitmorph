package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

type Profile struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	SSHKey   string `json:"sshKey,omitempty"`
	Default  bool   `json:"default"`
}

var profiles map[string]Profile
var configFile string

// cache of last-applied repo profile and its modtime (in-memory only)
var lastAppliedProfile string
var lastAppliedModTime time.Time

// Internal commands that should not trigger auto-switch
var skipAuto = map[string]bool{
	"activate":   true,
	"deactivate": true,
	"default":    true,
	"delete":     true,
	"edit":       true,
	"list":       true,
	"new":        true,
	"help":       true,
}

// RootCmd
var RootCmd = &cobra.Command{
	Use:                "gitmorph",
	Short:              "A CLI tool to seamlessly switch between Git identities",
	Long:               `GitMorph is a command-line tool that helps you manage and effortlessly switch between different Git accounts on your local machine.`,
	Args:               cobra.ArbitraryArgs,
	DisableFlagParsing: true, // pass all flags to Git
	Run: func(cmd *cobra.Command, args []string) {
		// If no arguments, show help and exit
		if len(args) == 0 {
			_ = cmd.Help()
			return
		}

		// Skip auto-switch for internal commands
		if !skipAuto[args[0]] {
			if err := applyRepoProfile(); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		}

		// Forward all args to Git
		if err := passthroughGit(args); err != nil {
			// passthroughGit already prints errors; exit with its code
			os.Exit(1)
		}
	},
}

func init() {
	homeDir, _ := os.UserHomeDir()
	configFile = filepath.Join(homeDir, ".gitmorph.json")

	RootCmd.CompletionOptions.DisableDefaultCmd = true

	loadProfiles()
}

// loadProfiles loads profiles from ~/.gitmorph.json
func loadProfiles() {
	profiles = make(map[string]Profile)

	data, err := os.ReadFile(configFile)
	if err == nil {
		json.Unmarshal(data, &profiles)
	}
}

// saveProfiles saves profiles to ~/.gitmorph.json
func saveProfiles() {
	data, _ := json.MarshalIndent(profiles, "", "  ")
	os.WriteFile(configFile, data, 0644)
}

// ------------------------
// Repo profile helpers
// ------------------------

// getRepoProfile reads ./gitmorph and returns (profileName, modTime, nil).
// If the file does not exist, returns ("", time.Time{}, os.ErrNotExist)
func getRepoProfile() (string, time.Time, error) {
	info, err := os.Stat(".gitmorph")
	if err != nil {
		return "", time.Time{}, err
	}

	data, err := os.ReadFile(".gitmorph")
	if err != nil {
		return "", time.Time{}, err
	}

	var m map[string]string
	if err := json.Unmarshal(data, &m); err != nil {
		return "", time.Time{}, err
	}

	if p, ok := m["profile"]; ok {
		return p, info.ModTime(), nil
	}
	return "", info.ModTime(), fmt.Errorf(".gitmorph missing 'profile' field")
}

// applyRepoProfile checks whether the repo .gitmorph changed since last applied.
// If changed (profile name or modtime), it applies the profile globally (silent).
// If .gitmorph is absent, it falls back to the default profile.
func applyRepoProfile() error {
	profileName, modTime, err := getRepoProfile()
	if err == nil {
		// .gitmorph exists
		if profileName == "" {
			return fmt.Errorf(".gitmorph has empty profile")
		}
		// Only apply if profile changed or file is newer
		if profileName != lastAppliedProfile || modTime.After(lastAppliedModTime) {
			if _, ok := profiles[profileName]; !ok {
				return fmt.Errorf("profile '%s' in .gitmorph not found in %s", profileName, configFile)
			}
			if err := switchToProfile(profileName, true); err != nil {
				return err
			}
			lastAppliedProfile = profileName
			lastAppliedModTime = modTime
		}
		return nil
	}

	// .gitmorph not present -> fallback to default
	var defaultName string
	for name, p := range profiles {
		if p.Default {
			defaultName = name
			break
		}
	}
	if defaultName == "" {
		return fmt.Errorf("no .gitmorph file found and no default profile set")
	}

	// If default differs from last applied, apply it
	if defaultName != lastAppliedProfile {
		if err := switchToProfile(defaultName, true); err != nil {
			return err
		}
		lastAppliedProfile = defaultName
		lastAppliedModTime = time.Time{} // no repo file
	}
	return nil
}

// passthroughGit forwards arguments to the system git command
func passthroughGit(args []string) error {
	if len(args) == 0 {
		return nil
	}

	gitPath, err := exec.LookPath("git")
	if err != nil {
		return fmt.Errorf("git not found in PATH")
	}

	cmd := exec.Command(gitPath, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
