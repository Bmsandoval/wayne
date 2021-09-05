package users

import (
	"github.com/bmsandoval/wayne/internal/db/models"
	"github.com/ztrue/tracerr"
)

var FindByUsernameSql = `
    SELECT id, sub, username, created_at, updated_at, deleted_at
    FROM users u
    WHERE username = ?
`

func (h *UserSvc) FindByUn(username string, user models.User) (*models.User, error) {
	statement, err := h.AppCtx.DB.Prepare(FindByUsernameSql)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	err = statement.QueryRow(username).
		Scan(&user.Id, &user.Sub, &user.Username, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil { return nil, tracerr.Wrap(err) }

	return &user, nil
}
