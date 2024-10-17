package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	numberOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		// field.Interface returns field's current value as an interface{}
		fn(val.String())
	case reflect.Struct:
		// val has a method NumField which returns the number of fields in the value.
		// This lets us iterate over the fields and call fn which passes our test.
		numberOfValues = val.NumField()
		getField = val.Field
	case reflect.Slice:
		numberOfValues = val.Len()
		getField = val.Index
	}

	for i := 0; i < numberOfValues; i++ {
		walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	// The reflect package has a function ValueOf which returns us a Value of a given
	// variable. This has ways for us to inspect a value, including its fields.
	val := reflect.ValueOf(x)

	// If val is of Kind Pointer we want to reassign val to the underlying value
	// before we can do NumField.
	if val.Kind() == reflect.Pointer {
		// We can get the underlying value with Elem()
		val = val.Elem()
	}

	return val
}
