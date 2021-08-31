package db
//
//import (
//	"database/sql"
//	"github.com/bmsandoval/wayne/db/models"
//	"github.com/square/squalor"
//	"reflect"
//)
//
//type Connection struct {
//	*sql.DB
//	SDB *squalor.DB
//	GreetingModel *squalor.Model
//}
//
//type bindable interface {
//	TableName() string
//	ModelName() string
//}
//
//var bindables = []bindable{
//	models.User{},
//}
//
//func BindModels(database *sql.DB) (*Connection, error) {
//	sdb, err := squalor.NewDB(database)
//	if err != nil {
//		return nil, err }
//
//	binding := &Connection{
//		DB:            database,
//		SDB:           sdb }
//
//	for _, bindable := range bindables {
//		model, err := sdb.BindModel(bindable.TableName(), bindable)
//		if err != nil {
//			return nil, err
//		} else {
//			SetField(binding, bindable.ModelName(), model)
//		}
//	}
//
//	return binding, nil
//}
//
//func SetField(conn *Connection, field string, value interface{}) {
//	v := reflect.ValueOf(conn).Elem().FieldByName(field)
//	if v.IsValid() {
//		v.Set(reflect.ValueOf(value))
//	}
//}