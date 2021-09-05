package DbContext

import "database/sql"

type Connection struct {
	Tx *sql.Tx
	*sql.DB
}

