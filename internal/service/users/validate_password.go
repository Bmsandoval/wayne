package users

import (
	"github.com/bmsandoval/wayne/db/models"
	"github.com/ztrue/tracerr"
	"golang.org/x/crypto/bcrypt"
)


var FindPasswordByUsernameSql = `
    SELECT id, sub, username, password, created_at, updated_at, deleted_at
    FROM users u
    WHERE username = ?
`

func (h *Helper) ValidatePassword(username string, password string, user models.User) (*models.User, error) {
	statement, err := h.AppCtx.DB.Prepare(FindPasswordByUsernameSql)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	err = statement.QueryRow(username).
		Scan(&user.Id, &user.Sub, &user.Username, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil { return nil, tracerr.Wrap(err) }

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, tracerr.Wrap(err)
	}
	return &user, nil
}
