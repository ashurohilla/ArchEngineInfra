package api

import "database/sql"

type APIServer struct {
	addr string
	db   *sql.db
}

func NewAPIServer(addr string, db *sql.DB) *ApIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}
