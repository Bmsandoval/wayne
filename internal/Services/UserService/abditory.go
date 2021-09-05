package UserService

// NOTE : Filename 'abditory' chosen to keep this file at the top of the list

import (
	"github.com/bmsandoval/wayne/internal/Db/models"
	"github.com/bmsandoval/wayne/pkg/AppContext"
)

// Bundlable adheres to Services.IBundlable and allows the service to be bundled
type Bundlable struct{}
// CreateService is a copy-pasta property
func(h Bundlable) CreateService(appCtx AppContext.Context) (interface{}, error) {
	return &UserSvc{
		AppCtx: appCtx,
	}, nil
}

// ServiceName Return the name of the service as it is defined in Services.Bundle
func (h Bundlable) ServiceName() string {
	return "UserSvc"
}



// UserSvc adheres to IUserSvc
type UserSvc struct {
	AppCtx AppContext.Context
}
type IUserSvc interface {
	Create(username string, password string) (string, error)
	FindByUn(username string, user models.User) (*models.User, error)
	FindBySub(sub string, user models.User) (*models.User, error)
	UsernameAvailable(username string) (bool, error)
	ValidatePassword(username string, password string, user models.User) (*models.User, error)
}
