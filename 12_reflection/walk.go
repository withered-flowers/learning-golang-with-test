package reflection

import "reflect"

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func walk(x interface{}, fn func(input string)) {
	// fn("Whatever the string is")

	// ? We will ask go to infer the value of the "any" given to x
	// val := reflect.ValueOf(x)

	// // ! This is not safe !
	// // ? We will try to look for the first field, if none, PANIC
	// field := val.Field(0)
	// // ? Return the value as string, this will wrong if the value is not string
	// fn(field.String())

	// Since we will have more than one field, loop it
	// for i := 0; i < val.NumField(); i++ {
	// 	field := val.Field(i)

	// 	// ! WRONG: This is for the assumption all field is string
	// 	// fn(field.String())

	// 	// ? RIGHT: We will make it call fn only if the field is string
	// 	if field.Kind() == reflect.String {
	// 		fn(field.String())
	// 	}

	// 	// ? What if the struct is nested?
	// 	// ? SOLUTION: use recursive
	// 	if field.Kind() == reflect.Struct {
	// 		// ? re-call the walk function, with field.Interface() as the "x" value
	// 		walk(field.Interface(), fn)
	// 	}
	// }

	// // ? Now we we need to handle pointer
	// if val.Kind() == reflect.Pointer {
	// 	// ? Extract underlying pointer value using Elem()
	// 	val = val.Elem()
	// }

	// // ? Refactor - using getValue
	// val := getValue(x)

	// ? Next Case - handling slice
	// if val.Kind() == reflect.Slice {
	// 	for i := 0; i < val.Len(); i++ {
	// 		// ? Solution: using recursive again
	// 		walk(val.Index(i).Interface(), fn)
	// 	}
	// 	// ! Don't forget to return
	// 	return
	// }

	// ? Refactor the code above, using switch case
	// for i := 0; i < val.NumField(); i++ {
	// 	field := val.Field(i)

	// 	switch field.Kind() {
	// 	case reflect.Struct:
	// 		walk(field.Interface(), fn)
	// 	case reflect.String:
	// 		fn(field.String())
	// 	}
	// }

	// // ? Refactor handle slice to inside switch case
	// val := getValue(x)

	// switch val.Kind() {
	// case reflect.Struct:
	// 	// ? If struct, loop using NumField()
	// 	for i := 0; i < val.NumField(); i++ {
	// 		// ? Fetch the value using val.Field, then recurse it
	// 		walk(val.Field(i).Interface(), fn)
	// 	}
	// case reflect.Slice:
	// 	// ? If slice, loop using Len()
	// 	for i := 0; i < val.Len(); i++ {
	// 		// ? Fetch the value using val.Index, then recurse it
	// 		walk(val.Index(i).Interface(), fn)
	// 	}
	// case reflect.String:
	// 	fn(val.String())
	// }

	// ? Next Refactor - reduce redundant recurse call
	val := getValue(x)

	numberOfValues := 0
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
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

// ? Refactor - fetchValue to it's own function
func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	// ? Now we we need to handle pointer
	if val.Kind() == reflect.Pointer {
		// ? Extract underlying pointer value using Elem()
		val = val.Elem()
	}

	return val
}
