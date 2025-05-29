# GitMorph

GitMorph is a powerful CLI tool that allows you to seamlessly switch between multiple Git identities on your local machine. Perfect for developers who work on different projects with various Git accounts.

<img width="826" alt="Screenshot 2025-05-25 at 9 59 25 AM" src="https://github.com/user-attachments/assets/c0801555-546a-4b69-a9b0-508d0b9c60ad" />

## Features

- Create and manage multiple Git profiles
- Easily switch between different Git identities (incl. per-profile SSH key)
- List all available profiles (shows SSH key path)
- Edit existing profiles
- Delete profiles
- Simple and intuitive command-line interface

## Installation

To install GitMorph, make sure you have Go installed on your system, then run:

```bash
go install github.com/abhigyan-mohanta/gitmorph@latest
````

### Update PATH

After installation, you may need to add the Go binaries directory to your system's `PATH` so you can run `gitmorph` from anywhere:

```bash
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc
source ~/.zshrc
```

## Usage

GitMorph provides the following commands:

### Create a new profile

```bash
gitmorph new
```

Prompts for:

* **Profile name**
* **Git username**
* **Git email**
* **SSH private key path** (leave blank for `~/.ssh/id_ed25519`)

### List all profiles

```bash
gitmorph list
```

Shows:

```
Available Git profiles:
- work     (Username: alice, Email: alice@corp.com, SSH: ~/.ssh/id_ed25519_work)
- personal (Username: alice123, Email: alice@gmail.com, SSH: ~/.ssh/id_ed25519)
```

### Switch to a profile

```bash
gitmorph switch <profile-name>
```

* Sets `user.name` and `user.email` globally
* Sets or unsets `core.sshCommand` to use the profile’s SSH key

### Edit a profile

```bash
gitmorph edit <profile-name>
```

Interactively update any of:

* Username
* Email
* SSH key path

Leave a prompt blank to keep the current value.

### Delete a profile

```bash
gitmorph delete <profile-name>
```

Removes the profile entry from `~/.gitmorph.json`.

## How It Works

GitMorph stores your Git profiles in a JSON file located at `~/.gitmorph.json`. Commands:

* `new` / `edit` / `delete` modify the JSON
* `switch` updates your global Git config:

   * `git config --global user.name <username>`
   * `git config --global user.email <email>`
   * `git config --global core.sshCommand "ssh -i <sshKey>"`

If you delete or switch back to a profile with no custom key, it unsets `core.sshCommand`.

## Code Structure

* `main.go`: Entry point
* `cmd/root.go`: Root command, loading/saving JSON
* `cmd/new.go`: `new` command
* `cmd/list.go`: `list` command
* `cmd/switch.go`: `switch` command
* `cmd/edit.go`: `edit` command
* `cmd/delete.go`: `delete` command

## SSH Configuration

You can still use `~/.ssh/config` if you like; GitMorph’s per-profile `core.sshCommand` will override it when set.

Example `~/.ssh/config`:

```plaintext
Host github.com
  HostName github.com
  User git
  IdentityFile ~/.ssh/id_ed25519

Host github.com-work
  HostName github.com
  User git
  IdentityFile ~/.ssh/id_ed25519_work
```

## Auto switch by .gitmorph config

1.  Add .gitmorph in root directory of git repo
    ```
    cat 'personal' >> .gitmorph
    ```
2. add script in bash profile, ie .bashrc/.zshrc or .profile, replace the DEFAULT_GITMORPH_PROFILE if need. 

```bash
# Define the default profile
DEFAULT_GITMORPH_PROFILE="work"

# Get the path of the original git command
ORIGINAL_GIT=$(command -v git)

# Define a function to automatically switch gitmorph profile
function check_and_switch_gitmorph() {
    local proceed_with_git=true  # Flag to determine if git command should proceed

    # Prevent recursive calls
    if [[ "$SKIP_GITMORPH" == "1" ]]; then
        return
    fi

    # Check if the current directory is a subdirectory of a git repository
    # Use the original git command to get the repository root directory
    if repo_root=$($ORIGINAL_GIT rev-parse --show-toplevel 2>/dev/null); then
        # Initialize profile_name as empty
        profile_name=""
        # Check if the .gitmorph file exists
        if [ -f "$repo_root/.gitmorph" ]; then
            # Read the profile name from the .gitmorph file
            profile_name=$(cat "$repo_root/.gitmorph")
            # Check if the profile name is empty
            if [ -z "$profile_name" ]; then
                echo ".gitmorph file is empty, using default profile: $DEFAULT_GITMORPH_PROFILE"
                profile_name=$DEFAULT_GITMORPH_PROFILE
            fi
        else
            echo "No .gitmorph file found in the repository root, using default profile: $DEFAULT_GITMORPH_PROFILE"
            profile_name=$DEFAULT_GITMORPH_PROFILE
        fi

        # Set the SKIP_GITMORPH environment variable to prevent recursive calls
        # Redirect both standard output and error output to a variable
        switch_output=$(SKIP_GITMORPH=1 gitmorph switch $profile_name 2>&1)
        success_pattern="Switched to Git profile"
        echo ">>>>$switch_output"

        if [[ "$switch_output" == *"$success_pattern"* ]]; then
            echo -e "Using gitmorph Profile: \033[31m$profile_name\033[0m"
        else
            # Switch failed, output relevant information and do not proceed with git command
            proceed_with_git=false
            echo -e "\033[31m==========Gitmorph Error============\033[0m"
            echo -e "\033[31mError: Failed to switch to gitmorph profile: $profile_name\033[0m"
            echo -e "\033[31mCheck the details below:\033[0m"
            echo "$switch_output"
            echo -e "\033[31mPlease ensure the profile exists and the SSH key is correctly configured.\033[0m"
            echo -e "\033[31mGit command aborted.\033[0m"
            echo -e "\033[31m==========End Gitmorph Error============\033[0m"
        fi
    fi

    # Export the flag to determine if git command should proceed
    if [[ "$proceed_with_git" == true ]]; then
        export GITMORPH_SWITCH_SUCCESSFUL=1
    else
        export GITMORPH_SWITCH_SUCCESSFUL=0
    fi
}

# Override the git command
function git() {
    check_and_switch_gitmorph

    # Check if git command should proceed
    if [[ "$GITMORPH_SWITCH_SUCCESSFUL" == "1" ]]; then
        $ORIGINAL_GIT "$@"
    else
        echo -e "\033[31mGit command aborted due to failed profile switch.\033[0m"
    fi
}
   ```

## Dependencies

```go
require (
    github.com/spf13/cobra v1.8.1
    github.com/spf13/pflag v1.0.5
)
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
