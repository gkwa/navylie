package navylie

import (
	"html/template"
	"log/slog"
	"os"

	"github.com/taylormonacelli/coalfoot"
)

func renderTemplate(tpl *coalfoot.TxtarTemplate, userSpecifiedDirectory string, templateData TemplateData) {
	tpl.FetchFromRemoteIfOld()

	tmpl, err := template.ParseFiles(tpl.LocalPathUnrendered)
	if err != nil {
		slog.Error("parseFiles", "path", tpl.LocalPathUnrendered, "error", err.Error())
		return
	}

	outputFile, err := os.Create(tpl.LocalPathRendered)
	if err != nil {
		slog.Error("create file", "path", tpl.LocalPathRendered, "error", err.Error())
		return
	}
	defer outputFile.Close()

	err = tmpl.Execute(outputFile, templateData)
	if err != nil {
		slog.Error("template execute", "path", tpl.LocalPathUnrendered, "error", err.Error())
		return
	}
}
