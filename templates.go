package navylie

import (
	"html/template"
	"log/slog"
	"os"
	"path/filepath"
)

func renderTemplate(outputDir string) {
	data := Data{
		ModuleName: filepath.Base(outputDir),
	}

	templatePath := "/tmp/navylie/templates/1.txtar"
	fname := "1-rendered.txtar"

	fetchTemplateToPath(templatePath)

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		slog.Error("parseFiles", "path", templatePath, "error", err.Error())
		return
	}

	outPath := filepath.Join(outputDir, fname)

	os.MkdirAll(outputDir, 0o755)

	outputFile, err := os.Create(outPath)
	if err != nil {
		slog.Error("create file", "path", outPath, "error", err.Error())
		return
	}
	defer outputFile.Close()

	err = tmpl.Execute(outputFile, data)
	if err != nil {
		slog.Error("template execute", "path", outPath, "error", err.Error())
		return
	}
}
