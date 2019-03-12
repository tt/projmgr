package db

import (
	"context"
	"net/http"

	"github.com/heroku/projmgr/db"
)

// For more information about context and why we're doing this,
// see https://blog.golang.org/context
type ctxkey int

var key ctxkey = 0

func WithDB(ctx context.Context, w http.ResponseWriter, r *http.Request) context.Context {
	c := db.NewClient()
	c.Open()

	go func() {
		<-ctx.Done()
		c.Close()
	}()

	return context.WithValue(ctx, key, c)
}

func FromContext(ctx context.Context) *db.Client {
	c, _ := ctx.Value(key).(*db.Client)
	return c
}
