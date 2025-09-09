package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

// switchToProfile switches Git config based on profile.
// If silent is true, suppresses printing and command output.
func switchToProfile(name string, silent bool) error {
	profile, exists := profiles[name]
	if !exists {
		return fmt.Errorf("profile '%s' does not exist", name)
	}

	// Determine whether to apply locally (repo) or globally
	local := false
	if data, err := os.ReadFile(".gitmorph"); err == nil {
		var f map[string]string
		if jsonErr := json.Unmarshal(data, &f); jsonErr == nil {
			if p, ok := f["profile"]; ok && p == name {
				local = true
			}
		}
	}

	var cmds []*exec.Cmd
	if local {
		// Apply to local repo
		cmds = []*exec.Cmd{
			exec.Command("git", "config", "--local", "user.name", profile.Username),
			exec.Command("git", "config", "--local", "user.email", profile.Email),
		}
	} else {
		// Apply to global
		cmds = []*exec.Cmd{
			exec.Command("git", "config", "--global", "user.name", profile.Username),
			exec.Command("git", "config", "--global", "user.email", profile.Email),
		}
	}

	run := func(c *exec.Cmd) error {
		if silent {
			devNull, err := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
			if err != nil {
				return err
			}
			defer devNull.Close()
			c.Stdout = devNull
			c.Stderr = devNull
		} else {
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr
		}
		return c.Run()
	}

	for _, c := range cmds {
		if err := run(c); err != nil {
			return fmt.Errorf("failed to set git config: %w", err)
		}
	}

	// Configure SSH key globally if present
	if profile.SSHKey != "" {
		sshCmd := fmt.Sprintf("ssh -i %s", profile.SSHKey)
		c := exec.Command("git", "config", "--global", "core.sshCommand", sshCmd)
		if err := run(c); err != nil {
			return fmt.Errorf("failed to set core.sshCommand: %w", err)
		}
	} else {
		// unset previous SSH key if exists
		c := exec.Command("git", "config", "--global", "--unset", "core.sshCommand")
		_ = run(c) // ignore errors
	}

	if !silent {
		scope := "global"
		if local {
			scope = "local"
		}
		fmt.Printf("Switched to profile '%s' (%s)\n", name, scope)
	}

	return nil
}
