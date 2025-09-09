# GitMorph 
**Now available on Homebrew as well**

GitMorph is a powerful CLI tool that allows you to seamlessly switch between multiple Git identities on your local machine. It ensures that all Git commands use the active profile’s identity and SSH key. Perfect for developers who work on different projects with various Git accounts and want consistent commits across repositories.

<img width="934" height="420" alt="Screenshot 2025-09-09 at 2 43 56 PM" src="https://github.com/user-attachments/assets/491d0b2c-8690-4f1b-bea3-b2388a1973b6" />

---

## ⚠️ Important Note: Git Command Wrapper

From **v3+**, GitMorph acts as a **wrapper for all Git commands** in repositories with an active GitMorph profile.

* Any Git command should now be prefixed with `gitmorph`.
* Examples:

```bash
gitmorph add .
gitmorph commit -m "Your commit message"
gitmorph push origin main
```

* This ensures the active profile’s Git identity and SSH key are correctly applied.
* Running plain `git` commands may bypass the active profile and cause mismatched commits.
* Older users please run gitmorph fix once after updating.

---

## Features

* Create and manage multiple Git profiles
* Easily switch between different Git identities (incl. per-profile SSH key)
* List all available profiles (shows SSH key path)
* Edit existing profiles
* Delete profiles
* Repo-specific auto-switch using `.gitmorph` file
* Simple and intuitive command-line interface

---

## Installation

Make sure you have **Go** installed, then run:

```bash
go install github.com/abhigyan-mohanta/gitmorph@latest
```

### Update PATH

After installation, add Go binaries to your `PATH`:

```bash
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc
source ~/.zshrc
```

---

## Usage

### Create a new profile

```bash
gitmorph new
```

Prompts for:

* **Profile name**
* **Git username**
* **Git email**
* **SSH private key path** (leave blank for `~/.ssh/id_ed25519`)
* **Set this profile as default?** (y/N)

---

### List all profiles

```bash
gitmorph list
```

Example output:

```
Available Git profiles:
NAME      USERNAME          EMAIL                     SSH KEY                     FLAGS
work      abhigyan6602      abhigyan@hostagedown.com  ~/.ssh/id_ed25519           [default]
personal  abhigyan-mohanta  underthunder02@gmail.com  ~/.ssh/id_ed25519_personal  [active]
```

---

### Set a default profile

```bash
gitmorph default <profile-name>
```

* Sets the specified profile as the global default.

Example:

```bash
gitmorph default personal
# Profile 'personal' is now the default.
```

---

### Activate a repo-specific profile

```bash
gitmorph activate <profile-name>
```

* Sets `user.name` and `user.email` globally
* Sets or unsets `core.sshCommand` to use the profile’s SSH key
* Adds a `.gitmorph` file in the project root to auto-use this profile

---

### Deactivate a repo-specific profile

```bash
gitmorph deactivate
```

* Deactivates the repo-specific profile
* Falls back to the global default profile

Example:

```bash
gitmorph deactivate
# Project-specific profile deactivated. Falling back to default profile.
# Switched to profile 'work' (global)
```

---

### Fix your GitMorph configuration

```bash
gitmorph fix
```

* Repairs `~/.gitmorph.json` if something went wrong
* Reapplies the default profile if needed

Example:

```bash
gitmorph fix
# Fixed ~/.gitmorph.json. 
# 'personal' is now the default profile.
# Change it using 'gitmorph default <profile>'
```

*Recommended for users migrating from v2 to v3+.*

---

### Edit a profile

```bash
gitmorph edit <profile-name>
```

Interactively update:

* Username
* Email
* SSH key path
* Default profile flag (current: false)

Leave a prompt blank to keep the current value.

---

### Delete a profile

```bash
gitmorph delete <profile-name>
```

Removes the profile entry from `~/.gitmorph.json`.

---

### Migration Note (v3+)

If upgrading from GitMorph v2:

* Run `gitmorph fix` to migrate your existing configuration
* Re-set your default profile using `gitmorph default <profile>` if needed
* Verify repo-specific profiles with `gitmorph deactivate`

---

## SSH Configuration

You can still use `~/.ssh/config`; GitMorph’s `core.sshCommand` overrides it when set.

Example:

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

---

## Dependencies

```go
require (
    github.com/spf13/cobra v1.8.1
    github.com/spf13/pflag v1.0.5
)
```

---

## Contributing

Contributions are welcome! Please submit a Pull Request.
