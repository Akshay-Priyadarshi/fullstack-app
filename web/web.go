package web

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

var (
	//go:embed dist/*
	dist embed.FS

	// Create sub filesystem stripping dist
	distFS = MustSubFS(dist, "dist")
)

func MustSubFS(originalFS embed.FS, strippedPath string) fs.FS {
	subFS, err := fs.Sub(originalFS, strippedPath)
	if err != nil {
		panic(err)
	}
	return subFS
}

func RegisterRoutes(app *fiber.App, path string) {
	app.Use(path, filesystem.New(filesystem.Config{
		Root:         http.FS(distFS),
		Index:        "index.html",
		NotFoundFile: "index.html",
	}))
}
