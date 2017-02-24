package app

import (
	"database/sql"
	"net/http"
)

type Context struct {
	DB *sql.DB
}

type Handler struct {
	*Context
	H func(*Context, http.ResponseWriter, *http.Request) (int, error)
}

func (ah Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status, err := ah.H(ah.Context, w, r)
	if err != nil {
		switch status {
		case http.StatusNotFound:
			http.NotFound(w, r)
		case http.StatusInternalServerError:
			http.Error(w, http.StatusText(status), status)
		default:
			http.Error(w, http.StatusText(status), status)
		}
	}
}
