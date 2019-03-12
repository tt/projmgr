package main

import (
	"context"
	"net/http"

	"github.com/guregu/kami"
	. "github.com/heroku/projmgr"
	"github.com/heroku/projmgr/http/middleware/db"
	"github.com/heroku/projmgr/mediators/projects"
)

func createProject(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	c := db.FromContext(ctx)

	project := &Project{}
	err := readJSON(r, project)
	if err != nil {
		writeError(w, err)
		return
	}

	m := &projects.Creator{ProjectService: c.ProjectService()}

	project, err = m.Call(project.Name)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, project)
}

func getProject(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	c := db.FromContext(ctx)

	id := ProjectID(kami.Param(ctx, "id"))

	project, err := c.ProjectService().Project(id)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, project)
}

func listProjects(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	c := db.FromContext(ctx)

	projects, err := c.ProjectService().Projects()
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, projects)
}
