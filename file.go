package navylie

import (
	"log/slog"
	"os"
)

func createFile(p string) *os.File {
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}

	return f
}

func closeFile(f *os.File) {
	err := f.Close()
	if err != nil {
		slog.Error("file close", "error", err.Error())
		os.Exit(1)
	}
}
