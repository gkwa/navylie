package navylie

import (
	"bytes"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
)

func runTxtar(txtarDir string) {
	txtarFname := "1-rendered.txtar"
	x := filepath.Join(txtarDir, txtarFname)

	var err error
	var txtarPath string

	txtarPath, err = filepath.Abs(x)
	if err != nil {
		slog.Error("error getting abs path", "file", txtarFname)
	}

	cmd := exec.Command("txtar-x", txtarPath)
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

	inputFile, err := os.Open(txtarPath)
	if err != nil {
		slog.Error("error opening input file", "path", txtarPath, "error", err.Error())
		return
	}
	defer inputFile.Close()

	if stdout.Len() > 0 {
		path := "output.txt"
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
		path := "error.txt"
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
