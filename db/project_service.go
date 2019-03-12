package db

import (
	"github.com/heroku/projmgr"
)

// CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
// CREATE TABLE projects (id UUID DEFAULT uuid_generate_v4(), name TEXT);

var _ projmgr.ProjectService = &projectService{}

type projectService struct {
	client *Client
}

func (s *projectService) Projects() ([]*projmgr.Project, error) {
	projects := []*projmgr.Project{}

	return projects, list(s.client.db, &projects, `SELECT * FROM projects;`)
}

func (s *projectService) Project(id projmgr.ProjectID) (*projmgr.Project, error) {
	project := &projmgr.Project{}

	return project, get(s.client.db, project, `SELECT * FROM projects WHERE id = $1;`, id)
}

func (s *projectService) CreateProject(project *projmgr.Project) error {
	tx := s.client.db.MustBegin()

	err := insert(tx, project, `INSERT INTO projects (name) VALUES (:name) RETURNING id;`, project)
	if err != nil {
		return err
	}

	return tx.Commit()
}
