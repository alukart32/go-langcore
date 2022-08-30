package controlflow

import (
	"fmt"
	"math/rand"
	"time"
)

// Examples of if-else control flow blocks
func IfElse() {
	ifElseBaseForm := `
if InitSimpleStatement; Condition {
	// do something
} else {
	// do something
}
`
	println("IfElse control base form:\n", ifElseBaseForm)
}

func IfElseAndRand() {
	rand.Seed(time.Now().UnixNano())

	if n := rand.Int(); n%2 == 0 {
		println(n, "is even number")
	} else {
		println(n, "is odd number")
	}

	n := rand.Int()
	if n%2 == 0 {
		println("Significand n: ", n, "is even number")
	} else {
		println("Significand n: ", n, "is odd number")
	}
}

func IfElseDayTimeRange() {
	if hour := time.Now().Local().Hour(); hour < 12 {
		println("Now is AM time.")
	} else if hour >= 19 {
		println("Now is PM time.")
	} else {
		println("Now is afternoon: ", hour, "h")
	}
}

func ForBaseForm() {
	forBaseForm := `
for InitSimpleStatement; Condition; PostSimpleStatement {
	// do something
}`

	println("\nFor control base form:", forBaseForm)
}

func ForCountDown() {
	println("\nCountDown#start")
	for i := 0; i < 3; i++ {
		println("count-i: ", i)
		// The left i is a new declared variable,
		// and the right i is the loop variable.
		i := i
		// The new declared variable is modified, but
		// the old one (the loop variable) is not yet.
		i = 10
		println("local-for-new-i: ", i)
		_ = i
	}
}

func ForBreak() {
	println("\nFor break")
	i := 0
	for {
		if i >= 10 {
			println("i == ", i, " has to break out control flow block")
			break
		}
		println("i: ", i)
		i++
	}
}

func ForContinue() {
	println("\nFor continue")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	println()
}

func SwitchBaseForm() {
	switchBaseForm := `
switch InitSimpleStatement; CompareOperand0 {
case CompareOperandList1:
	// do something
case CompareOperandList2:
	// do something
...
case CompareOperandListN:
	// do something
default:
	// do something
}`
	println("\nSwitch control block base form:\n", switchBaseForm)
}

func SwitchModN() {
	println("\nSwitch mod 9")

	rand.Seed(time.Now().UnixNano())
	switch n := rand.Intn(100); n % 9 {
	case 0:
		println(n, "is a multiple of 9.")

		// Different from many other languages,
		// in Go, the execution will automatically
		// jumps out of the switch-case block at
		// the end of each branch block.
		// No "break" statement is needed here.
	case 1, 2, 3:
		println(n, "mod 9 is 1, 2 or 3.")
		break // this "break" statement is nonsense.
	case 4, 5, 6:
		fmt.Println(n, "mod 9 is 4, 5 or 6.")
	default:
		fmt.Println(n, "mod 9 is 7 or 8.")
	}
}

func SwitchFallthrough() {
	println("\nSwitch mod 9")

	rand.Seed(time.Now().UnixNano())
	switch n := rand.Intn(100); n % 9 {
	case 0:
		println(n, "is a multiple of 9.")

		// Different from many other languages,
		// in Go, the execution will automatically
		// jumps out of the switch-case block at
		// the end of each branch block.
		// No "break" statement is needed here.
		fallthrough
	case 1, 2, 3:
		println(n, "mod 9 is 1, 2 or 3.")
		fallthrough
	case 4, 5, 6:
		fmt.Println(n, "mod 9 is 4, 5 or 6.")
		fallthrough
	default:
		fmt.Println("n =", n)
	}
}

func GoToInc() {
	println("\nGoto inc value")

	i := 0
NextVal:
	println("i: ", i)
	if i++; i < 5 {
		goto NextVal
	}
}

func FuzzyGoTo() {
Start:
	i := 0
Next:
	if i >= 5 {
		// error: jumps over declaration of k
		// goto Exit jumps over variable declaration at line 183
		// goto Exit
		// ok
		goto Start
	}

	k := i + i
	fmt.Println(k)
	i++
	goto Next

	// This label is declared in the scope of k,
	// but its use is outside of the scope of k.
	//Exit:
}

func GoTofindSmallestPrimeLargerThan() {
	println("\nFind smallest prime larger than value")
	for i := 90; i < 100; i++ {
		n := findSmallestPrimeLargerThan(i)
		fmt.Print("The smallest prime number larger than ")
		fmt.Println(i, "is", n)
	}
}

func findSmallestPrimeLargerThan(n int) int {
Outer:
	for n++; ; n++ {
		for i := 2; ; i++ {
			switch {
			case i*i > n:
				break Outer
			case n%i == 0:
				continue Outer
			}
		}
	}
	return n
}
