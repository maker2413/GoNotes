package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		// field.Interface returns field's current value as an interface{}
		fn(val.String())
	case reflect.Struct:
		// val has a method NumField which returns the number of fields in the value.
		// This lets us iterate over the fields and call fn which passes our test.
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	// A case can have multiple triggers
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	case reflect.Chan:
		for {
			// Recv receives and returns a value from the channel v. It panics if v's Kind
			// is not Chan. The receive blocks until a value is ready. The boolean value ok
			// is true if the value x corresponds to a send on the channel, false if it is a
			// zero value received because the channel is closed.
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}
	case reflect.Func:
		// Call calls the function v with the input arguments in.
		// func (v Value) Call(in []Value) []Value
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
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
