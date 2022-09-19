package types

import (
	"fmt"
)

func InterfaceTypeAssertion1() {
	var i interface{} = 123
	n, ok := i.(int)
	fmt.Printf("Dynamic type (123) is int?\n\t%d %v\n", n, ok)

	f, ok := i.(float64)
	fmt.Printf("Dynamic type (123) is float64?\n\t%f %v\n", f, ok)
}

type Writer interface {
	Write(buf []byte) (int, error)
}

type DummyWriter struct{}

func (DummyWriter) Write(buf []byte) (int, error) {
	return len(buf), nil
}

func InterfaceTypeAssertion2() {
	var x interface{} = DummyWriter{} // dynamic type is DummyWriter
	var y interface{} = "123"         // dynamic type is String

	var w Writer
	var ok bool

	w, ok = x.(Writer)
	fmt.Printf("Dynamic type (DummyWriter) is Writer?\n\t%v %v\n", w, ok)
	x, ok = w.(interface{})
	fmt.Printf("Dynamic type (DummyWriter) is interface{}?\n\t%v %v\n", x, ok)

	w, ok = y.(Writer)
	fmt.Printf("Dynamic type (string) is Writer?\n\t%v %v\n", w, ok)
}

func InterfceTypeSwitch() {
	i := []interface{}{
		1, "123", []int{4, 5, 6}, map[string]int{"city": 1},
	}

	for _, el := range i {
		switch v := el.(type) {
		case nil:
			fmt.Printf("nil\n")
		case int, int32, int64, float32, float64:
			fmt.Printf("number: %v\n", v)
		case string:
			fmt.Printf("string: %v\n", v)
		case []int:
			fmt.Printf("int slice: %v\n", v)
		default:
			fmt.Printf("unknown type: %T %v\n", v, v)
		}
	}
}
