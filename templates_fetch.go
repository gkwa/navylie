package navylie

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
)

func fetchTemplateToPath(templatePath string) {
	url := "https://raw.githubusercontent.com/taylormonacelli/navylie/master/templates/1.txtar"

	directory := filepath.Dir(templatePath)

	if err := os.MkdirAll(directory, os.ModePerm); err != nil {
		fmt.Printf("Error creating directories: %v\n", err)
		return
	}

	fileName := filepath.Join(directory, filepath.Base(url))

	if _, err := os.Stat(fileName); !os.IsNotExist(err) {
		x, _ := filepath.Abs(fileName)
		slog.Debug("file already exists, not refetching", "path", x)
	} else {
		response, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error fetching the file: %v\n", err)
			return
		}
		defer response.Body.Close()

		file, err := os.Create(fileName)
		if err != nil {
			fmt.Printf("Error creating the file: %v\n", err)
			return
		}
		defer file.Close()

		_, err = io.Copy(file, response.Body)
		if err != nil {
			fmt.Printf("Error saving the file: %v\n", err)
			return
		}

		x, _ := filepath.Abs(fileName)
		slog.Debug("file saved", "path", x)
	}
}
