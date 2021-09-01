package users

import (
	"github.com/bmsandoval/wayne/db/models"
	"github.com/ztrue/tracerr"
)

var FindBySubSql = `
    SELECT id, sub, username, created_at, updated_at, deleted_at
    FROM users u
    WHERE Id = UNHEX( REPLACE( ?,'-','' ))
`

func (h *Helper) FindBySub(sub string, user models.User) (*models.User, error) {
	statement, err := h.AppCtx.DB.Prepare(FindBySubSql)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	err = statement.QueryRow(sub).
		Scan(&user.Id, &user.Sub, &user.Username, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil { return nil, tracerr.Wrap(err) }

	return &user, nil
}
