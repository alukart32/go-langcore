package other

import (
	"fmt"
)

func TrySomeFuzzyStuff() {
	a := float64(5)
	b := 8.12
	fmt.Printf("a + b = %v", a+b)

	var i interface{}
	if i != nil {
		println("not nil")
	}
	println("nil")
}

func f1(a, b int) {
	{
		a, b := a, b
		println(a, b)
	}
}
