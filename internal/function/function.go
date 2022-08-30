package function

func F() (x int, y bool) {
	println(x, y) // 0 false
	return
}

func AnonymousFunc() {
	// This anonymous function has no parameters
	// but has two results.
	x, y := func() (int, int) {
		println("This function has no parameters.")
		return 3, 4
	}() // Call it. No arguments are needed.

	// The following anonymous function have no results.

	func(a, b int) {
		// The following line prints: a*a + b*b = 25
		println("a*a + b*b =", a*a+b*b)
	}(x, y) // pass argument x and y to parameter a and b.

	func(x int) {
		// The parameter x shadows the outer x.
		// The following line prints: x*x + y*y = 32
		println("x*x + y*y =", x*x+y*y)
	}(y) // pass argument y to parameter x.

	func() {
		// The following line prints: x*x + y*y = 25
		println("x*x + y*y =", x*x+y*y)
	}() // no arguments are needed.
}

func F2(int, string, _ string) string {
	println("F2")
	return "string"
}

func Sum(numbers ...int64) (sum int64) {
	// The type of values is []int64.
	sum = 0
	for _, v := range numbers {
		sum += v
	}
	return
}

func add(a, b int) int {
	return a + b
}

func FuncValue() {
	var addFunc func(int, int) int = add
	println("addFunc(5, 6)~add(5, 6) = ", addFunc(5, 6))

	sumGt20 := func(sum int) bool {
		return sum > 20
	}
	println("anonymous func sumGt20(addFunc(5, 19)): 5 + 19 = ", sumGt20(addFunc(5, 19)))

	// This function returns a function (a closure).
	isMultipleOfX := func(x int) func(n int) bool {
		count := 0
		return func(n int) bool {
			count++
			println("isMultipleOfX calls: ", count)
			return n%x == 0
		}
	}
	var isMultipleOf3 = isMultipleOfX(3)
	println("assign anonymous functions to function variable: isMultipleOf3 8 = ", isMultipleOf3(8))
	println("assign anonymous functions to function variable: isMultipleOf3 13 = ", isMultipleOf3(13))
	println("assign anonymous functions to function variable: isMultipleOf3 12 = ", isMultipleOf3(12))
	println("assign anonymous functions to function variable: isMultipleOf3 44 = ", isMultipleOf3(44))

}
