package types

import (
	"fmt"
)

type Filter interface {
	About() string
	Process([]int) []int
}

// UniqueFilter is used to remove duplicate numbers.
type UniqueFilter struct{}

func (UniqueFilter) About() string {
	return "remove duplicate numbers"
}

func (UniqueFilter) Process(inputs []int) []int {
	outs := make([]int, 0, len(inputs))
	pusheds := make(map[int]bool)
	for _, n := range inputs {
		if !pusheds[n] {
			pusheds[n] = true
			outs = append(outs, n)
		}
	}
	return outs
}

// MultipleFilter is used to keep only
// the numbers which are multiples of
// the MultipleFilter as an int value.
type MultipleFilter int

func (mf MultipleFilter) About() string {
	return fmt.Sprintf("keep multiples of %v", mf)
}

func (mf MultipleFilter) Process(inputs []int) []int {
	outs := make([]int, 0, len(inputs))
	for _, n := range inputs {
		if n%int(mf) == 0 {
			outs = append(outs, n)
		}
	}
	return outs
}

// With the help of polymorphism, only one
// "filterAndPrint" function is needed.
func filter(filter Filter, inputs []int) []int {
	// Calling the methods of "filter" will call the
	// methods of the dynamic value in interface value.
	filtered := filter.Process(inputs)
	fmt.Printf("%s :\n\t%v\n", filter.About(), filtered)
	return filtered
}

func Polymorphism() {
	numbers := []int{1, -3, 4, 67, -90, 4, 25, 1, 22}
	fmt.Println("before filtering:\n\t", numbers)

	// Three non-interface values are boxed into
	// three Filter interface slice element values.
	filters := []Filter{
		UniqueFilter{},
		MultipleFilter(2),
		MultipleFilter(5),
	}
	for _, f := range filters {
		numbers = filter(f, numbers)
	}
}
