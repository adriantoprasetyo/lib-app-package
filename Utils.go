package libapppackage

import (
	"reflect"
	"strings"
)

// StructToMap skips zero-value fields
func StructToMap(input interface{}) map[string]interface{} {
	out := make(map[string]interface{})
	v := reflect.ValueOf(input)
	t := reflect.TypeOf(input)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		// Get GORM tag or fallback to field name
		tag := field.Tag.Get("gorm")
		name := strings.Split(tag, ";")[0]
		if name == "" || name == "-" {
			name = field.Name
		}

		// Only include if value is not zero
		if !isZeroValue(value) {
			out[name] = value.Interface()
		}
	}

	return out
}

// isZeroValue checks if a reflect.Value is the zero value for its type
func isZeroValue(v reflect.Value) bool {
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}
