package users

import (
	"github.com/ztrue/tracerr"
	"golang.org/x/crypto/bcrypt"
	"log"
)

var ValidatePasswordSql = `
    SELECT sub
    FROM users u
    WHERE username = $1 AND password = $2
`

func (h *Helper) ValidatePassword(username string, password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	statement, err := h.AppCtx.DB.Prepare(ValidatePasswordSql)
	if err != nil {
		return "", tracerr.Wrap(err)
	}

	rows, err := statement.Query(username, hashedPassword)
	if err != nil { return "", err }
	uSub := ""
	for rows.Next() {
		if err := rows.Scan(&uSub); err != nil {
			return "", err
		}
		break
	}


	return uSub, nil
}
