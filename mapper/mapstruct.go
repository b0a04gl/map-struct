package mapper

import (
	"errors"
	"fmt"
	"reflect"
)


type Person struct {
	Name    string
	Age     int
	Address string
}


func MapToStruct(data map[string]interface{}, target interface{}) error {
	targetValue := reflect.ValueOf(target)

	if err := validateTarget(targetValue); err != nil {
		return err
	}

	targetElem := targetValue.Elem()
	targetType := targetElem.Type()

	for i := 0; i < targetElem.NumField(); i++ {
		field := targetElem.Field(i)
		fieldName := targetType.Field(i).Name

		if value, ok := data[fieldName]; ok {
			if field.CanSet() {
				fieldValue := reflect.ValueOf(value)
				if fieldValue.Type().AssignableTo(field.Type()) {
					field.Set(fieldValue)
				} else {
					return fmt.Errorf("type mismatch for field %s", fieldName)
				}
			}
		}
	}

	return nil
}


func validateTarget(target reflect.Value) error {
	if target.Kind() != reflect.Ptr || target.IsNil() {
		return errors.New("target must be a non-nil pointer to a struct")
	}
	return nil
}
