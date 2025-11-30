package env

import (
	"os"
	"strings"
)

func ReadENV(path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	for line := range strings.SplitSeq(string(file), "\n") {
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		if err := os.Setenv(key, value); err != nil {
			return err
		}
	}

	return nil
}
