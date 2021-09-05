package service

import (
	"github.com/bmsandoval/wayne/internal/service/users"
	"github.com/bmsandoval/wayne/internal/utilities/appcontext"
	"reflect"
)

type Bundle struct {
	UserSvc users.IUserSvc
}

var bundlableServices = []IBundlable{
	users.Bundlable{},
}

type IBundlable interface {
	CreateService(appCtx appcontext.Context) (interface{}, error)
	ServiceName() string
}

func NewBundle(appCtx appcontext.Context) (*Bundle, error) {
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