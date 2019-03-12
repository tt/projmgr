package projmgr

type ProjectID string

type Project struct {
	ID   ProjectID `db:"id" json:"id"`
	Name string    `db:"name" json:"name"`
}

// Client creates a connection to the services.
type Client interface {
	ProjectService() ProjectService
}

// ProjectService represents a service for managing projects.
type ProjectService interface {
	CreateProject(project *Project) error
	Project(id ProjectID) (*Project, error)
	Projects() ([]*Project, error)
}
