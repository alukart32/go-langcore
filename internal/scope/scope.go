package scope

import (
	"fmt"
)

const y = 789

var x int = 123

func Scopes() {
	// The x variable shadows the above declared
	// package-level variable x.
	var x = true

	// A nested code block.
	{
		// Here, the left x and y are both
		// new declared variable. The right
		// ones are declared in outer blocks.
		x, y := x, y

		// In this code block, the just new
		// declared x and y shadow the outer
		// declared same-name identifiers.
		x, z := !x, y/10 // only z is new declared
		y /= 100

		{
			x, y, z := x, y, z
			fmt.Printf("Deep code block ~ x: %v, y: %v, z: %v\n", !x, y+3, z+2)
		}

		fmt.Printf("Middle code block ~ x: %v, y: %v, z: %v\n", x, y, z) // false 7 78
	}
	println(x) // true
	//println(z) // error: z is undefined.
}
