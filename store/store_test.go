package store_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/nicola-strappazzon/password-manager/store"
)

func TestListAndItems(t *testing.T) {
	root := t.TempDir()
	for _, name := range []string{"personal/site.gpg", "personal/nested/other.gpg", "root.gpg", ".git/hidden.gpg", "personal/.hidden.gpg", "personal/notes.txt"} {
		path := filepath.Join(root, name)
		if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte("opaque ciphertext"), 0600); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.Symlink(t.TempDir(), filepath.Join(root, "outside")); err != nil {
		t.Fatal(err)
	}
	entries, err := store.List(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 5 {
		t.Fatalf("unexpected visible entries: %d", len(entries))
	}
	items, err := store.Items(root, "personal")
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 1 || items[0].Path != "personal/site" {
		t.Fatalf("unexpected items: %+v", items)
	}
	for _, path := range []string{"../", "/tmp", ".git", "personal/../personal", "outside"} {
		if _, err := store.Items(root, path); err == nil {
			t.Errorf("accepted unsafe directory %q", path)
		}
	}
}
