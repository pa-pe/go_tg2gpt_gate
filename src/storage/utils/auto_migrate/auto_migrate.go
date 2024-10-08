package auto_migrate

import (
	"fmt"
	"gorm.io/gorm"
	"reflect"
)

var registeredModels []interface{}

// RegisterModel make a list for db.AutoMigrate in auto_migrate.Execute()
func RegisterModel(model interface{}) {
	registeredModels = append(registeredModels, model)
}

// Execute do db.AutoMigrate for all registered models with auto_migrate.RegisterModel(model)
func Execute(db *gorm.DB) error {
	for _, model := range registeredModels {
		if checkContainsGormModel(model) {
			fmt.Printf("Migrating model: %T\n", model)
			if err := db.AutoMigrate(model); err != nil {
				return err
			}
		} else {
			fmt.Printf("Skipping model: %T (does not contain gorm.Model)\n", model)
		}
	}
	return nil
}

// checking is model containing gorm.Model field
func checkContainsGormModel(model interface{}) bool {
	v := reflect.ValueOf(model)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()

	// all fields recursion
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Type == reflect.TypeOf(gorm.Model{}) {
			return true
		}
	}
	return false
}
