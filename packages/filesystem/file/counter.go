package file

import (
	"fmt"
	"github.com/greeflas/itea_golang/filesystem"
	"os"
)

func CountFiles(path string) (int, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, fmt.Errorf("CountFiles: cannot open dir: %q: %w", path, err)
	}

	var filesCount int

	for _, entry := range entries {
		if !entry.IsDir() {
			filesCount++

			continue
		}

		info, err := entry.Info()
		if err != nil {
			return 0, fmt.Errorf("CountFiles: cannot get dir info: %w", err)
		}

		count, err := CountFiles(filesystem.BuildPath(path, info.Name()))
		if err != nil {
			return 0, err
		}

		filesCount += count
	}

	return filesCount, nil
}
