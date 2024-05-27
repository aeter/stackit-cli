package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

func TestWrite(t *testing.T) {
	tests := []struct {
		description  string
		folderName   string
		folderExists bool
	}{
		{
			description: "write config file",
			folderName:  "",
		},
		{
			description: "write config file to new folder",
			folderName:  "new-folder",
		},
		{
			description:  "write config file to existing folder",
			folderName:   "existing-folder",
			folderExists: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			configPath := filepath.Join(os.TempDir(), tt.folderName, "config.json")
			viper.SetConfigFile(configPath)
			folderPath = filepath.Dir(configPath)

			if tt.folderExists {
				err := os.MkdirAll(folderPath, os.ModePerm)
				if err != nil {
					t.Fatalf("expected error to be nil, got %v", err)
				}
			}

			err := Write()
			if err != nil {
				t.Fatalf("expected error to be nil, got %v", err)
			}

			// Check if the file was created
			_, err = os.Stat(configPath)
			if os.IsNotExist(err) {
				t.Fatalf("expected file to exist, got %v", err)
			}

			// Delete the file
			err = os.Remove(configPath)
			if err != nil {
				t.Fatalf("expected error to be nil, got %v", err)
			}

			// Delete the folder
			if tt.folderName != "" {
				err = os.Remove(folderPath)
				if err != nil {
					t.Fatalf("expected error to be nil, got %v", err)
				}
			}
		})
	}
}