package Services

import (
	"github.com/bmsandoval/wayne/internal/Services/UserService"
	"github.com/bmsandoval/wayne/pkg/AppContext"
	"reflect"
)

type Bundle struct {
	UserSvc UserService.IUserSvc
}

var bundlableServices = []IBundlable{
	UserService.Bundlable{},
}

type IBundlable interface {
	CreateService(appCtx AppContext.Context) (interface{}, error)
	ServiceName() string
}

func NewBundle(appCtx AppContext.Context) (*Bundle, error) {
	bundle := &Bundle{}

	for _, bundlableService := range bundlableServices {
		helper, err := bundlableService.CreateService(appCtx)
		if err != nil {
			return nil, err
		}
		SetField(bundle, bundlableService.ServiceName(), helper)
	}

	return bundle, nil
}

func SetField(bundler *Bundle, field string, value interface{}) {
	v := reflect.ValueOf(bundler).Elem().FieldByName(field)
	if v.IsValid() {
		v.Set(reflect.ValueOf(value))
	}
}