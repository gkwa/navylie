package navylie

import (
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/taylormonacelli/ashpalm"
	"github.com/taylormonacelli/coalfoot"
)

func Main(userProjectDir string) int {
	tpl := coalfoot.NewTxtarTemplate()
	tpl.FetchFromRemoteIfOld()

	slog.Debug("user project dir", "dir", userProjectDir)

	userProjectDirAbs, err := filepath.Abs(userProjectDir)
	if err != nil {
		slog.Error("filepath.abs", "path", userProjectDirAbs, "error", err.Error())
	}

	templateData := TemplateData{
		ModuleName:     filepath.Base(userProjectDirAbs),
		GithubUsername: "taylormonacelli",
	}

	renderTemplate(tpl, userProjectDir, templateData)

	err = tpl.Extract(userProjectDirAbs)
	if err != nil {
		slog.Error("extracting", "error", err.Error())
		return 1
	}

	slog.Debug("running func", "func", "runGoModTidy")
	runGoModTidy(userProjectDirAbs)

	return 0
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
