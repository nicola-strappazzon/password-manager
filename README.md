# Password Manager

[![Test](https://github.com/nicola-strappazzon/password-manager/actions/workflows/test.yaml/badge.svg?branch=main)](https://github.com/nicola-strappazzon/password-manager/actions/workflows/test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/nicola-strappazzon/password-manager)](https://goreportcard.com/report/github.com/nicola-strappazzon/password-manager)

This is another Unix-style password manager written in Go to save your data with OpenPGP encryption.

## Install on macOS

Using [Homebrew](https://brew.sh/):

```bash
brew install nicola-strappazzon/tap/password-manager
```

## Install using go

If you have Go installed, you can install the password-manager binary like this:

```bash
go install github.com/nicola-strappazzon/password-manager@latest
```

The binary will be placed in your `GOBIN` directory, which defaults to `~/go/bin`. Depending on how Go is installed, this directory may or may not be in your `PATH`.

## Install on Linux - DEB

This method works on Debian, Ubuntu, and related distributions. The example below is for Intel/AMD systems. For ARM architectures, replace "amd64" with "arm64" in both commands.

```bash
curl -LO https://github.com/nicola-strappazzon/password-manager/releases/latest/download/password-manager_amd64.deb
sudo apt install ./password-manager_amd64.deb
```

## Install on Linux - RPM

This method works on RHEL, CentOS, Fedora, Amazon Linux, SUSE, Mandriva, Oracle Linux, and similar distributions. The example below is for Intel/AMD systems. For ARM architectures, replace "amd64" with "arm64" in both commands.

```bash
curl -LO https://github.com/nicola-strappazzon/password-manager/releases/latest/download/password-manager_amd64.rpm
sudo rpm -Uvh password-manager_amd64.rpm
```

## Install on Windows

A Windows build is currently unavailable, but the Linux build works fine under WSL.

> [!WARNING]
> This project is under active development and may be unstable. Use at your own risk.
