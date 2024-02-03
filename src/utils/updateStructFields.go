package utils

import (
	"reflect"
)

//Original is the struct we want to update, while incoming is struct that contains the new data we want to update with
func UpdateStructFields(original interface{}, incoming interface{}) {
	originalValue := reflect.ValueOf(original).Elem()
	originalType := originalValue.Type()
	incomingValue := reflect.ValueOf(incoming).Elem()

	for i := 0; i < originalValue.NumField(); i++ {
		originalField := originalValue.Field(i)
		fieldName := originalType.Field(i).Name
		incomingField := incomingValue.FieldByName(fieldName)

		if originalField.Kind() == reflect.Struct && incomingField.Kind() == reflect.Struct {
			UpdateStructFields(originalField.Addr().Interface(), incomingField.Addr().Interface())
		} else if !isEmpty(incomingField) && incomingField.IsValid() {
			originalField.Set(incomingField)
		}

	}
}
func isEmpty(field reflect.Value) bool {
    switch field.Kind() {
    case reflect.String:
        return field.String() == ""
    case reflect.Bool:
		//doesn't work well for bool especially when the incoming field
		//is false because it will read false as being empty. better to 
		//use ptr for bool fields
        return !field.Bool()
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        return field.Int() == 0
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
        return field.Uint() == 0
    case reflect.Float32, reflect.Float64:
        return field.Float() == 0.0
    case reflect.Complex64, reflect.Complex128:
        return field.Complex() == 0+0i
    case reflect.Array, reflect.Slice, reflect.Map:
        return field.Len() == 0
    case reflect.Ptr, reflect.Interface:
        return field.IsNil()
    default:
        return false
    }
}