package cmd

import (
	"os"
	"testing"
)

func TestCommandLineExecution(t *testing.T){
	t.Run("parse directory argument", func(t *testing.T) {
		want := "path/to/directory"
		os.Args = []string{"hamocha", want}

		dir, err := ParseDirectory()

		if err != nil {
			t.Errorf("expected no error but got: %v", err)
		}

		if dir != want {
			t.Errorf("expected %s but got %s", want, dir)
		}
	})
	t.Run("return directory entries", func(t *testing.T) {
		want := "path/to/directory"
		os.Args = []string{"hamocha", want}

		entries, err := GetDirectoryEntries(want)

		if err != nil {
			t.Errorf("expected no error but got: %v", err)
		}

		if len(entries) == 0 {
			t.Errorf("expected some entries but got none")
		}
	})

	t.Run("parse directory argument with no argument", func(t *testing.T) {
		os.Args = []string{"hamocha"}

		_, err := ParseDirectory()

		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("category file by extension", func(t *testing.T) {
		
	})
}
