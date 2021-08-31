package users

import (
	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
	"golang.org/x/crypto/bcrypt"
)

var UserInsertSql = `
INSERT INTO users(username, password,id)
VALUES ( ?,?,UNHEX( REPLACE( ?,'-','' )));
`

func (h *Helper) Create(username string, password string) (string, error) {
	uid := uuid.New()

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", tracerr.Wrap(err)
	}

	statement, err := h.AppCtx.DB.Prepare(UserInsertSql)
	if err != nil {
		return "", tracerr.Wrap(err)
	}

	if _, err := statement.Exec(username, hash, uid.String()); err != nil {
		return "", tracerr.Wrap(err)
	}

	return uid.String(), nil
}
