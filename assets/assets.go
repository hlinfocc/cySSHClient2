package assets

import (
	"embed"
	"net/http"
)

var (
	//go:embed dist/*
	content embed.FS

	FileSystem http.FileSystem
)

func init() {
	FileSystem = http.FS(content)
}

func GetIndexHtml() ([]byte, error) {
	indexHTML, err := content.ReadFile("dist/index.html")
	if err != nil {
		// c.String(http.StatusInternalServerError, "Failed to read index.html: %v", err)
		return nil, err
	}
	return indexHTML, nil
}
