package funk

import (
	"fmt"
	"reflect"
)

// ForEachOption defines the options for ForEach
type ForEachOption struct {
	Reverse bool
}

var boolType = reflect.TypeOf(true)

// ForEach iterates over elements of collection and invokes iteratee for each element.
func ForEach(arr interface{}, predicate interface{}, options ...ForEachOption) {
	if !IsIteratee(arr) {
		panic("First parameter must be an iteratee")
	}

	option := ForEachOption{}
	if len(options) > 0 {
		option = options[0]
	}

	var (
		funcValue = reflect.ValueOf(predicate)
		arrValue  = reflect.ValueOf(arr)
		arrType   = arrValue.Type()
		funcType  = funcValue.Type()
		numIn     = funcType.NumIn()
		numOut    = funcType.NumOut()
	)

	if numOut == 1 {
		if t := funcType.Out(0); !t.ConvertibleTo(boolType) {
			panic("Map function's return is not compatible with bool.")
		}
	}

	if arrType.Kind() == reflect.Slice || arrType.Kind() == reflect.Array {
		if !IsFunc(predicate, []int{1, 2}, []int{0, 1}) {
			panic("Second argument must be a function with one/two parameter")
		}

		inOffset := IfInt(numIn != 2, 0, 1)
		// Checking whether element type is convertible to function's first argument's type.
		if t := arrValue.Type().Elem(); !t.ConvertibleTo(funcType.In(inOffset)) {
			panic("Map function's argument is not compatible with type of array.")
		}

		switch {
		case numIn == 1 && !option.Reverse:
			for i := 0; i < arrValue.Len(); i++ {
				outs := funcValue.Call([]reflect.Value{arrValue.Index(i)})
				if numOut == 1 && !outs[0].Convert(boolType).Interface().(bool) {
					break
				}
			}
		case numIn == 2 && !option.Reverse:
			for i := 0; i < arrValue.Len(); i++ {
				outs := funcValue.Call([]reflect.Value{reflect.ValueOf(i), arrValue.Index(i)})
				if numOut == 1 && !outs[0].Convert(boolType).Interface().(bool) {
					break
				}
			}
		case numIn == 1 && option.Reverse:
			for i := arrValue.Len() - 1; i >= 0; i-- {
				outs := funcValue.Call([]reflect.Value{arrValue.Index(i)})
				if numOut == 1 && !outs[0].Convert(boolType).Interface().(bool) {
					break
				}
			}
		case numIn == 2 && option.Reverse:
			for i := arrValue.Len() - 1; i >= 0; i-- {
				outs := funcValue.Call([]reflect.Value{reflect.ValueOf(i), arrValue.Index(i)})
				if numOut == 1 && !outs[0].Convert(boolType).Interface().(bool) {
					break
				}
			}
		}
	}

	if arrType.Kind() == reflect.Map {
		if !IsFunc(predicate, []int{2, 3}, nil) {
			panic("Second argument must be a function with two/three parameters")
		}

		// Type checking for Map<key, value> = (key, value)
		inOffset := IfInt(numIn != 3, 0, 1)

		if t := arrType.Key(); !t.ConvertibleTo(funcType.In(inOffset)) {
			panic(fmt.Sprintf("function first argument is not compatible with %v", t))
		}

		if t := arrType.Elem(); !t.ConvertibleTo(funcType.In(inOffset + 1)) {
			panic(fmt.Sprintf("function second argument is not compatible with %v", t))
		}

		switch numIn {
		case 2:
			for _, key := range arrValue.MapKeys() {
				outs := funcValue.Call([]reflect.Value{key, arrValue.MapIndex(key)})
				if numOut == 1 && !outs[0].Convert(boolType).Interface().(bool) {
					break
				}
			}
		default: // 3
			for i, key := range arrValue.MapKeys() {
				outs := funcValue.Call([]reflect.Value{reflect.ValueOf(i), key, arrValue.MapIndex(key)})
				if numOut == 1 && !outs[0].Convert(boolType).Interface().(bool) {
					break
				}
			}
		}
	}
}

// ForEachRight iterates over elements of collection from the right and invokes iteratee for each element.
func ForEachRight(arr interface{}, predicate interface{}) {
	ForEach(arr, predicate, ForEachOption{Reverse: true})
}

// Head gets the first element of array.
func Head(arr interface{}) interface{} {
	value := redirectValue(reflect.ValueOf(arr))
	valueType := value.Type()

	kind := value.Kind()

	if kind == reflect.Array || kind == reflect.Slice {
		if value.Len() == 0 {
			return nil
		}

		return value.Index(0).Interface()
	}

	panic(fmt.Sprintf("Type %s is not supported by Head", valueType.String()))
}

// Last gets the last element of array.
func Last(arr interface{}) interface{} {
	value := redirectValue(reflect.ValueOf(arr))
	valueType := value.Type()

	kind := value.Kind()

	if kind == reflect.Array || kind == reflect.Slice {
		if value.Len() == 0 {
			return nil
		}

		return value.Index(value.Len() - 1).Interface()
	}

	panic(fmt.Sprintf("Type %s is not supported by Last", valueType.String()))
}

// Initial gets all but the last element of array.
func Initial(arr interface{}) interface{} {
	value := redirectValue(reflect.ValueOf(arr))
	valueType := value.Type()

	kind := value.Kind()

	if kind == reflect.Array || kind == reflect.Slice {
		length := value.Len()

		if length <= 1 {
			return arr
		}

		return value.Slice(0, length-1).Interface()
	}

	panic(fmt.Sprintf("Type %s is not supported by Initial", valueType.String()))
}

// Tail gets all but the first element of array.
func Tail(arr interface{}) interface{} {
	value := redirectValue(reflect.ValueOf(arr))
	valueType := value.Type()

	kind := value.Kind()

	if kind == reflect.Array || kind == reflect.Slice {
		length := value.Len()

		if length <= 1 {
			return arr
		}

		return value.Slice(1, length).Interface()
	}

	panic(fmt.Sprintf("Type %s is not supported by Initial", valueType.String()))
}
