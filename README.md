# pm - Password Manager

[![Test](https://github.com/nicola-strappazzon/password-manager/actions/workflows/test.yaml/badge.svg?branch=main)](https://github.com/nicola-strappazzon/password-manager/actions/workflows/test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/nicola-strappazzon/password-manager)](https://goreportcard.com/report/github.com/nicola-strappazzon/password-manager)

A Unix-style password manager written in Go that stores data encrypted with OpenPGP.

> [!WARNING]
> This project is under active development and may be unstable. Use at your own risk.

## Install

Using [Homebrew](https://brew.sh/):

```bash
brew install nicola-strappazzon/tap/password-manager
```

Using Go:

```bash
go install github.com/nicola-strappazzon/password-manager@latest
```

The binary will be placed in your `GOBIN` directory, which defaults to `~/go/bin`. Make sure it is in your `PATH`.

## Setup

Run the setup command to configure the application for first use:

```bash
pm setup
```

It will guide you through generating or importing an OpenPGP key pair and will print the environment variables you need to add to your shell profile:

```bash
export PM_PUBLICKEY="$HOME/.password-manager/public.asc"
export PM_PRIVATEKEY="$HOME/.password-manager/private.asc"
```

## Commands

| Command | Description |
|---------|-------------|
| `pm ls` | List all items in tree format |
| `pm add <path>` | Add or update an encrypted item |
| `pm show <path>` | Decrypt and show an item |
| `pm edit <path>` | Edit an encrypted item |
| `pm remove <path>` | Remove an encrypted item |
| `pm otp <path>` | Generate an OTP code |
| `pm generate` | Generate a random password |
| `pm file <path>` | Manage files inside an encrypted item |
| `pm setup` | Configure the application for first use |

## Usage

```bash
# List all items
pm ls

# Add a new item
pm add personal/github -f username -v john
pm add personal/github -f password

# Show the password
pm show personal/github

# Show all fields
pm show personal/github -a

# Copy the password to clipboard
pm show personal/github -c

# Generate an OTP code
pm otp personal/github

# Generate a random password
pm generate

# Remove an item
pm remove personal/github
```
