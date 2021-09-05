package users

// NOTE : Filename 'abditory' chosen to keep this file at the top of the list

import (
	"github.com/bmsandoval/wayne/internal/db/models"
	"github.com/bmsandoval/wayne/internal/utilities/appcontext"
)



// Bundlable adheres to service.IBundlable and allows the service to be bundled
type Bundlable struct{}
// CreateService is a copy-pasta property
func(h Bundlable) CreateService(appCtx appcontext.Context) (interface{}, error) {
	return &UserSvc{
		AppCtx: appCtx,
	}, nil
}
// ServiceName Return the name of the service as it is defined in service.Bundle
func (h Bundlable) ServiceName() string {
	return "UserSvc"
}



// UserSvc adheres to IUserSvc
type UserSvc struct {
	AppCtx appcontext.Context
}
type IUserSvc interface {
	Create(username string, password string) (string, error)
	FindByUn(username string, user models.User) (*models.User, error)
	FindBySub(sub string, user models.User) (*models.User, error)
	UsernameAvailable(username string) (bool, error)
	ValidatePassword(username string, password string, user models.User) (*models.User, error)
}
