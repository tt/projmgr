package projects

import (
	"github.com/heroku/projmgr"
)

type Creator struct {
	ProjectService projmgr.ProjectService
}

func (c *Creator) Call(name string) (*projmgr.Project, error) {
	project := &projmgr.Project{Name: name}

	// TODO: perform validation

	return project, c.ProjectService.CreateProject(project)
}
