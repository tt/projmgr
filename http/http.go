package main

import (
	"github.com/guregu/kami"
	"github.com/heroku/projmgr/http/middleware/db"
)

func main() {
	kami.Cancel = true

	kami.Use("/", db.WithDB)

	kami.Post("/projects", createProject)
	kami.Get("/projects/:id", getProject)
	kami.Get("/projects", listProjects)

	kami.Serve()
}
