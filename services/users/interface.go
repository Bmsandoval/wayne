package users

import (
	"github.com/bmsandoval/wayne/library/appcontext"
)

type Helper struct {
	AppCtx appcontext.Context
}
type Helpable struct{}

func(h Helpable) NewHelper(appCtx appcontext.Context) (interface{}, error) {
	return &Helper{
		AppCtx: appCtx,
	}, nil
}

func (h Helpable) ServiceName() string {
	return "UserSvc"
}

type Service interface {
	Create(username string, password string) (string, error)
	ValidatePassword(username string, password string) (bool, error)
}
