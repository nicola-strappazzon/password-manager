# Password Manager

[![Test](https://github.com/nicola-strappazzon/password-manager/actions/workflows/test.yaml/badge.svg?branch=main)](https://github.com/nicola-strappazzon/password-manager/actions/workflows/test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/nicola-strappazzon/password-manager)](https://goreportcard.com/report/github.com/nicola-strappazzon/password-manager)

This is another Unix-style password manager written in Go to save your data with OpenPGP encryption.

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

The binary will be placed in your `GOBIN` directory, which defaults to `~/go/bin`. Depending on how Go is installed, this directory may or may not be in your `PATH`.

## Setup

Run the setup command to configure the application for first use:

```bash
pm setup
```

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
