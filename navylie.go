package navylie

import (
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/taylormonacelli/ashpalm"
	"github.com/taylormonacelli/coalfoot"
)

func Main(userProjectDir string) {
	x := coalfoot.NewTxtarTemplate()
	x.FetchIfNotFound()

	slog.Debug("user project dir", "dir", userProjectDir)

	userProjectDirAbs, err := filepath.Abs(userProjectDir)
	if err != nil {
		slog.Error("filepath.abs", "path", userProjectDirAbs, "error", err.Error())
	}

	userModuleName := filepath.Base(userProjectDirAbs)
	renderTemplate(x, userProjectDir, userModuleName)

	slog.Debug("running func", "func", "runTxtar")
	runTxtar(x.LocalPathRendered, userProjectDirAbs)

	slog.Debug("running func", "func", "runGoModTidy")
	runGoModTidy(userProjectDirAbs)
}

func runGoModTidy(txtarDir string) {
	cmd := exec.Command("go", "mod", "tidy")
	cwd := txtarDir

	code, outStr, errStr := ashpalm.RunCmd(cmd, cwd)
	if code != 0 {
		slog.Error("runcmd", "cmd", cmd.String(), "stdout", outStr, "stderr", errStr, "code", code)
		os.Exit(1)
	}

	slog.Debug("runcmd", "cmd", cmd.String(), "stdout", outStr, "stderr", errStr, "code", code)
}

func runTxtar(txtarPath, userProjectDir string) {
	cmd := exec.Command("txtar-x", "-C", userProjectDir, txtarPath)

	cwd := filepath.Dir(txtarPath)
	code, outStr, errStr := ashpalm.RunCmd(cmd, cwd)

	if code != 0 {
		slog.Error("runcmd", "cmd", cmd.String(), "stdout", outStr, "stderr", errStr, "code", code)
		os.Exit(1)
	}

	slog.Debug("runcmd", "cmd", cmd.String(), "stdout", outStr, "stderr", errStr, "code", code)
}
