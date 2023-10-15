package navylie

import (
	"bytes"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
)

func Main(outDir string) {
	renderTemplate(outDir)

	y, err := filepath.Abs(outDir)
	if err != nil {
		slog.Error("stuff", "path", y)
	}

	txtarDir := y
	runTxtar(txtarDir)
	runGoModTidy(txtarDir)
}

func runGoModTidy(txtarDir string) {
	var err error

	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = txtarDir

	slog.Debug("running command", "cmd", cmd.String(), "cwd", txtarDir)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		slog.Error("error running command", "error", err.Error())
	}

	outStr, errStr := stdout.String(), stderr.String()

	if stdout.Len() > 0 {
		path := "output-tidy.txt"
		outputFile, err := os.Create(path)
		if err != nil {
			slog.Error("error creating output file", "path", path, "error", err.Error())
			return
		}
		defer outputFile.Close()

		_, err = outputFile.WriteString(outStr)
		if err != nil {
			slog.Error("error writing to output file", "path", path, "error", err.Error())
			return
		}

	}

	if stderr.Len() > 0 {
		path := "error-tidy.txt"
		errorFile, err := os.Create(path)
		if err != nil {
			slog.Error("error creating error file", "path", path, "error", err.Error())
			return
		}
		defer errorFile.Close()

		_, err = errorFile.WriteString(errStr)
		if err != nil {
			slog.Error("Error writing to error file: %v\n", err)
			return
		}
	}
}
