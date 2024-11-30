package utils

import (
	"os"
	"path/filepath"
)

func GetDBPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".pwm/database.db"), nil
}
