package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func FindConfigFile(env string, maxAttempts int) (string, error) {
	filename := fmt.Sprintf("cfg/%s.config.yaml", env)
	for _ = range maxAttempts {
		fmt.Println("searching for config file", filename)
		if _, err := os.Stat(filename); !os.IsNotExist(err) {
			return filename, nil
		}
		filename = filepath.Join("..", filename)
	}
	return "", fmt.Errorf("config file not found after %d attempts", maxAttempts)
}
