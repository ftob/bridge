package statistics

import (
	"database/sql"
	"github.com/ftob/bridge"
)

type repository struct {
	conn *sql.Conn
}

func New(conn *sql.Conn) bridge.Repository {
	return &repository{conn:conn}
}

func Store(stat *bridge.Stat) {
	
}