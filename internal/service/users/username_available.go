package users

import (
	"github.com/ztrue/tracerr"
)

var UsernameAvailableSql = `
	SELECT COUNT(*) = 0
	FROM (
		SELECT 1
		FROM users u
		WHERE username = ?
	    LIMIT 1
	) AS found;
`

func (h *UserSvc) UsernameAvailable(username string) (bool, error) {
	statement, err := h.AppCtx.DB.Prepare(UsernameAvailableSql)
	if err != nil {
		return false, tracerr.Wrap(err)
	}

	available := false
	err = statement.QueryRow(username).
		Scan(&available)
	if err != nil { return false, tracerr.Wrap(err) }

	return available, nil
}
