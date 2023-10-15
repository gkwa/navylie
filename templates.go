package navylie

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
)

func renderTemplate(outputDir string) {
	data := Data{
		ModuleName: filepath.Base(outputDir),
	}

	templatePath := "/tmp/navylie/templates/1.txtar"

	fetchTemplateToPath(templatePath)

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatal(err)
	}

	fname := "1-rendered.txtar"
	outPath := filepath.Join(outputDir, fname)

	os.MkdirAll(outputDir, 0o755)

	outputFile, err := os.Create(outPath)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	err = tmpl.Execute(outputFile, data)
	if err != nil {
		log.Fatal(err)
	}
}
