package users

import (
	"github.com/google/uuid"
	"github.com/ztrue/tracerr"
	"golang.org/x/crypto/bcrypt"
)

var UserInsertSql = `
INSERT INTO users(username,password,id)
VALUES ( ?,?,UNHEX( REPLACE( ?,'-','' )));
`


func (h *Helper) Create(username string, password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	uid := uuid.New()

	statement, err := h.AppCtx.DB.Prepare(UserInsertSql)
	if err != nil {
		return "", tracerr.Wrap(err)
	}

	if _, err := statement.Exec(username, hash, uid ); err != nil {
		return "", tracerr.Wrap(err)
	}

	return uid.String(), nil
}
