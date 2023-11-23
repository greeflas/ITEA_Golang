package directory

import (
	"fmt"
	"github.com/greeflas/itea_golang/filesystem"
	"os"
)

func CountDirs(path string) (int, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, fmt.Errorf("CountDirs: cannot open dir: %q: %w", path, err)
	}

	var dirsCount int

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			return 0, fmt.Errorf("CountDirs: cannot get dir info: %w", err)
		}

		dirsCount++

		count, err := CountDirs(filesystem.BuildPath(path, info.Name()))
		if err != nil {
			return 0, err
		}

		dirsCount += count
	}

	return dirsCount, nil
}
