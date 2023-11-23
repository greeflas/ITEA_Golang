package main

import "os"

func initialize() {
	dirs := []string{"./test/a/b", "./test/c"}
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			panic(err)
		}
	}

	files := []string{"./test/a/b/e.txt", "./test/d.txt"}
	for _, file := range files {
		f, err := os.Create(file)
		if err != nil {
			panic(err)
		}

		if err := f.Close(); err != nil {
			panic(err)
		}
	}
}
