# GitMorph

GitMorph is a powerful CLI tool that allows you to seamlessly switch between multiple Git identities on your local machine. Perfect for developers who work on different projects with various Git accounts.

<img width="929" alt="Screenshot 2024-10-12 at 2 22 09 AM" src="https://github.com/user-attachments/assets/4860e7cb-0a2d-4ffc-bc31-f6f60b3d4b75">

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

## Dependencies

```go
require (
    github.com/spf13/cobra v1.8.1
    github.com/spf13/pflag v1.0.5
)
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.