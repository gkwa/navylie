package navylie

import (
	"bytes"
	"log/slog"
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
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = txtarDir

	stdErrLog := "error-tidy.txt"
	stdOutLog := "output-tidy.txt"

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	slog.Debug("running command", "cmd", cmd.String(), "cwd", txtarDir)
	err := cmd.Run()
	if err != nil {
		slog.Error("error running command", "error", err.Error())
	}

	outStr, errStr := stdout.String(), stderr.String()

	if stdout.Len() > 0 {
		f := createFile(stdOutLog)
		defer closeFile(f)

		f.WriteString(outStr)
	}

	if stderr.Len() > 0 {
		f := createFile(stdErrLog)
		defer closeFile(f)

		f.WriteString(errStr)
	}
}
