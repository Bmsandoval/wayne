package configs

import (
	"github.com/spf13/viper"
	"reflect"
)

type DatabaseConfiguration struct {
	DbMigrationLocation string `json:"MIGRATION_LOCATION"`
	DbSchema string `json:"MYSQL_SCHEMA"`
	DbHost string `json:"MYSQL_HOST"`
	DbPort string `json:"MYSQL_PORT"`
}

type DatabaseSecrets struct {
	DbUser string `json:"MYSQL_USER"`
	DbPass string `json:"MYSQL_PASS"`
}

func GetDatabaseConfig(vipe viper.Viper) DatabaseConfiguration {
	var newDatabaseConfiguration DatabaseConfiguration
	t := reflect.TypeOf(newDatabaseConfiguration)

	for i := 0; i < t.NumField(); i++ {
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := t.Field(i)

		if field.Name == "Secrets" { continue }

		// Get the field tag value
		tag := field.Tag.Get("json")

		if tag == "" { continue }
		v := reflect.ValueOf(&newDatabaseConfiguration).Elem().FieldByName(field.Name)
		if v.IsValid() {
			tagValue := vipe.GetString(tag)
			v.Set(reflect.ValueOf(tagValue))
		}
	}

	return newDatabaseConfiguration
}
