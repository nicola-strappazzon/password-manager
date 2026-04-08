package git

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/nicola-strappazzon/password-manager/internal/config"
)

type StatusEntry struct {
	Code string
	File string
}

func (e StatusEntry) Label() string {
	switch {
	case e.Code == "??":
		return "untracked"
	case strings.HasPrefix(strings.TrimSpace(e.Code), "A"):
		return "new file"
	case strings.Contains(e.Code, "D"):
		return "deleted"
	default:
		return "modified"
	}
}

func storePath() string {
	return config.GetPath("")
}

func IsRepo() bool {
	return exec.Command("git", "-C", storePath(), "rev-parse", "--is-inside-work-tree").Run() == nil
}

func Branch() string {
	out, err := exec.Command("git", "-C", storePath(), "branch", "--show-current").Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}

func StatusEntries() []StatusEntry {
	out, err := exec.Command("git", "-C", storePath(), "status", "--porcelain").Output()
	if err != nil {
		return nil
	}

	var entries []StatusEntry
	for _, line := range strings.Split(string(out), "\n") {
		if len(line) < 3 {
			continue
		}
		entries = append(entries, StatusEntry{
			Code: line[:2],
			File: strings.TrimSpace(line[3:]),
		})
	}
	return entries
}

func UnpushedCommits() []string {
	out, err := exec.Command("git", "-C", storePath(), "log", "@{u}..HEAD", "--oneline").Output()
	if err != nil {
		return nil
	}

	var commits []string
	for _, line := range strings.Split(strings.TrimSpace(string(out)), "\n") {
		if line != "" {
			commits = append(commits, line)
		}
	}
	return commits
}

func UncommittedChangesWarning() string {
	if !IsRepo() {
		return ""
	}

	n := len(StatusEntries())
	if n == 0 {
		return ""
	}

	return fmt.Sprintf("Warning: %d uncommitted change(s) in the password store.", n)
}

func UnpushedCommitsWarning() string {
	if !IsRepo() {
		return ""
	}

	n := len(UnpushedCommits())
	if n == 0 {
		return ""
	}

	return fmt.Sprintf("Warning: %d unpushed commit(s) in the password store.", n)
}
