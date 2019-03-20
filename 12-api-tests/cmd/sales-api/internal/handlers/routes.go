package handlers

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/ardanlabs/service-training/12-api-tests/internal/platform/web"
)

// API constructs an http.Handler with all application routes defined.
func API(db *sqlx.DB, log *log.Logger) http.Handler {

	app := web.New(log)

	p := Products{db: db, log: log}

	app.Handle(http.MethodPost, "/v1/products", p.Create)
	app.Handle(http.MethodGet, "/v1/products", p.List)
	app.Handle(http.MethodGet, "/v1/products/{id}", p.Get)

	return app
}
