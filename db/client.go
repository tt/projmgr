package db

import (
	"github.com/heroku/projmgr"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var _ projmgr.Client = &Client{}

type Client struct {
	// Services
	projectService projectService

	db *sqlx.DB
}

func NewClient() *Client {
	c := &Client{}
	c.projectService.client = c
	return c
}

func (c *Client) Open() error {
	db, err := sqlx.Connect("postgres", "dbname=lyra_test sslmode=disable")
	if err != nil {
		return err
	}

	c.db = db
	return nil
}

func (c *Client) Close() error {
	if c.db != nil {
		return c.db.Close()
	}

	return nil
}

// ProjectService returns the project service associated with the client.
func (c *Client) ProjectService() projmgr.ProjectService { return &c.projectService }
