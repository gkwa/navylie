package navylie

import (
	"log/slog"
	"os/exec"
	"path/filepath"

	"github.com/taylormonacelli/ashpalm"
)

func Main(outDir string) {
	renderTemplate(outDir)

	txtarDir, err := filepath.Abs(outDir)
	if err != nil {
		slog.Error("filepath.abs", "path", txtarDir, "error", err.Error())
	}

	runTxtar(txtarDir)
	runGoModTidy(txtarDir)
}

func runGoModTidy(txtarDir string) {
	cmd := exec.Command("go", "mod", "tidy")
	cwd := txtarDir

	code, outStr, errStr := ashpalm.RunCmd(cmd, cwd)
	slog.Debug("runcmd", "cmd", cmd.String(), "stdout", outStr, "stderr", errStr, "code", code)
}

func runTxtar(txtarDir string) {
	txtarPath := filepath.Join(txtarDir, "1-rendered.txtar")
	cmd := exec.Command("txtar-x", txtarPath)

	cwd := txtarDir
	code, outStr, errStr := ashpalm.RunCmd(cmd, cwd)
	slog.Debug("runcmd", "cmd", cmd.String(), "stdout", outStr, "stderr", errStr, "code", code)
}
