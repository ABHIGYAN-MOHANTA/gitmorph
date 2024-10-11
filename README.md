# GitMorph

GitMorph is a powerful CLI tool that allows you to seamlessly switch between multiple Git identities on your local machine. Perfect for developers who work on different projects with various Git accounts.

## Features

- Create and manage multiple Git profiles
- Easily switch between different Git identities
- List all available profiles
- Simple and intuitive command-line interface

## Installation

To install GitMorph, make sure you have Go installed on your system, then run:

```bash
go get github.com/abhigyan-mohanta/gitmorph
```

## Usage

GitMorph provides the following commands:

### Create a new profile

```bash
gitmorph new
```

This command will prompt you to enter a profile name, Git username, and Git email.

### Switch to a profile

```bash
gitmorph switch <profile-name>
```

This command switches your global Git configuration to the specified profile.

### List all profiles

```bash
gitmorph list
```

This command displays all available Git profiles.

## How It Works

GitMorph stores your Git profiles in a JSON file located at `~/.gitmorph.json`. When you switch profiles, it updates your global Git configuration using the `git config --global` command.

## Code Structure

The project is structured as follows:

- `main.go`: Entry point of the application
- `cmd/root.go`: Defines the root command and common functionality
- `cmd/new.go`: Implements the "new" command to create profiles
- `cmd/switch.go`: Implements the "switch" command to change profiles
- `cmd/list.go`: Implements the "list" command to display profiles

## SSH Configuration

For seamless Git operations with multiple accounts, you can set up SSH configurations. Below is an example of how your SSH config file (`~/.ssh/config`) should look:

```plaintext
# Work GitHub Account
Host github.com
    HostName github.com
    User git
    IdentityFile ~/.ssh/id_ed25519
    AddKeysToAgent yes
    
# Personal GitHub Account
Host github.com-personal
    HostName github.com
    User git
    IdentityFile ~/.ssh/id_ed25519_personal
    AddKeysToAgent yes
```

### Note on SSH Key Generation

To create an SSH key for your personal account, you can use the following command:

```bash
ssh-keygen -t ed25519 -C "your_email@example.com" -f ~/.ssh/id_ed25519_personal
```

### Usage Tips

- When using the main `id_ed25519.pub`, the normal command for pushing changes works as expected:

```bash
git push -u origin main
```

- For other accounts, you need to specify the SSH command:

```bash
GIT_SSH_COMMAND="ssh -i ~/.ssh/id_ed25519_personal" git push -u origin main
```

- Alternatively, you can configure the SSH command to use your SSH config file:

```bash
GIT_SSH_COMMAND="ssh -F ~/.ssh/config" git push -u origin main
```

## Dependencies

GitMorph uses the following external libraries:

```go
require (
    github.com/inconshreveable/mousetrap v1.1.0 // indirect
    github.com/spf13/cobra v1.8.1 // indirect
    github.com/spf13/pflag v1.0.5 // indirect
)
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
