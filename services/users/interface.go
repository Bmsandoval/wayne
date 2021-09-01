package users

import (
	"github.com/bmsandoval/wayne/db/models"
	"github.com/bmsandoval/wayne/library/appcontext"
)

// Helper One helper for each helpable
type Helper struct {
	AppCtx appcontext.Context
}

// Helpable service functions are branches off this object
type Helpable struct{}
// NewHelper copy-pasta property
func(h Helpable) NewHelper(appCtx appcontext.Context) (interface{}, error) {
	return &Helper{
		AppCtx: appCtx,
	}, nil
}
// ServiceName Return the name of the service as it is defined in the bundle.go
func (h Helpable) ServiceName() string {
	return "UserSvc"
}

// Service the interface your helpable (and therefore your service functions) must adhere to
// ^ if this doesn't match the function you will get an error like 'reflect.Set: value of type *users.Helper is not assignable to type users.Service'
type Service interface {
	Create(username string, password string) (string, error)
	FindByUn(username string, user models.User) (*models.User, error)
	FindBySub(sub string, user models.User) (*models.User, error)
	UsernameAvailable(username string) (bool, error)
	ValidatePassword(username string, password string, user models.User) (*models.User, error)
}
