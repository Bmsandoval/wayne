package service

import (
	"github.com/bmsandoval/wayne/library/appcontext"
	"github.com/bmsandoval/wayne/services/users"
	"reflect"
)

type Bundle struct {
	UserSvc users.Service
}

var bundlables = []bundlable{
	users.Helpable{},
}

type bundlable interface {
	NewHelper(appCtx appcontext.Context) (interface{}, error)
	ServiceName() string
}

func NewBundle(appCtx appcontext.Context) (*Bundle, error) {
	bundle := &Bundle{}

	for _, bundlable := range bundlables {
		helper, err := bundlable.NewHelper(appCtx)
		if err != nil {
			return nil, err
		}
		SetField(bundle, bundlable.ServiceName(), helper)
	}

	return bundle, nil
}

func SetField(bundler *Bundle, field string, value interface{}) {
	v := reflect.ValueOf(bundler).Elem().FieldByName(field)
	if v.IsValid() {
		v.Set(reflect.ValueOf(value))
	}
}