package users

import (
	"github.com/ztrue/tracerr"
)

var UsernameAvailableSql = `
    SELECT TOP 1 1
    FROM users u
    WHERE username = ?
`

func (h *Helper) UsernameAvailable(username string) (bool, error) {
	statement, err := h.AppCtx.DB.Prepare(FindByUsernameSql)
	if err != nil {
		return false, tracerr.Wrap(err)
	}

	available := false
	err = statement.QueryRow(username).
		Scan(&available)
	if err != nil { return false, tracerr.Wrap(err) }

	return available, nil
}
