package types

import (
	"encoding/json"
	"fmt"
	"reflect"
)

const Size = 5

func ContainerTypeDeclaration() {
	/* Array types */
	var string5 [Size]string = [Size]string{}
	println("\nArray type declaration\n")
	fmt.Printf("var string5 [Size]string = [Size]string{}, type = %v\n", reflect.TypeOf(string5))

	string5Short := [Size]string{}
	fmt.Printf("string5Short := [Size]string{}, type = %v\n", reflect.TypeOf(string5Short))

	// Element type is a slice type: []int
	arrSlice := [Size][]int{}
	fmt.Printf("arrSlice := [Size][]int{}, type = %v\n", reflect.TypeOf(arrSlice))

	/* Slice types *

	[]bool
	[]int64
	// Element type is a map type: map[int]bool
	[]map[int]bool
	// Element type is a pointer type: *int
	[]*int

	*/

	/* Map types */
	println("\nMap type declaration\n")
	map1 := map[string]int{}
	fmt.Printf("map1 := map[string]int{}, type = %v\n", reflect.TypeOf(map1))

	mapOfArray := map[string][Size]int{}
	fmt.Printf("mapOfArray := map[string][Size]int{}, type = %v\n", reflect.TypeOf(mapOfArray))

	mapOfSlices := map[bool][]string{}
	fmt.Printf("mapOfSlices := map[bool][]string{}, type = %v\n", reflect.TypeOf(mapOfSlices))

	type (
		Value = int
		Key   = string
	)
	customTypedMap := map[Key]Value{}
	fmt.Printf("customTypedMap := map[Key]Value{}, type = %v\n", reflect.TypeOf(customTypedMap))
}

func ArraySliceCompositeliterals() {
	// The followings slice composite literals
	// are equivalent to each other.
	println("\nSlice composite literals::")

	printStringSliceType := func(name string, s []string) {
		fmt.Printf("%s : %v, type %v\n", name, s, reflect.TypeOf(s))
	}

	slice1 := []string{"break", "continue", "fallthrough"}
	slice2 := []string{0: "break", 1: "continue", 2: "fallthrough"}
	slice3 := []string{2: "fallthrough", 1: "continue", 0: "break"}
	slice4 := []string{2: "fallthrough", 0: "break", "continue"}
	printStringSliceType("slice1", slice1)
	printStringSliceType("slice2", slice2)
	printStringSliceType("slice3", slice3)
	printStringSliceType("slice4", slice4)

	slice5 := []map[string]int{{"C": 1972}, {"Python": 1991}, {"Go": 2009}}
	fmt.Printf("slice5 : %v, type %v\n", slice5, reflect.TypeOf(slice5))

	slice6 := []int{}
	fmt.Printf("slice6 : %v, type %v\n", slice6, reflect.TypeOf(slice6))

	slice7 := []int(nil)
	fmt.Printf("slice7 : %v (nil), type %v\n", slice7, reflect.TypeOf(slice7))

	var a [5]int

	// The followings array composite literals
	// are equivalent to each other.
	println("\nAarray composite literals::", len(a))

	printBoolSliceType := func(name string, s [4]bool) {
		fmt.Printf("%s : %v, type %v\n", name, s, reflect.TypeOf(s))
	}

	arr1 := [4]bool{false, true, true, false}
	arr2 := [4]bool{0: false, 1: true, 2: true, 3: false}
	arr3 := [4]bool{1: true, true}
	arr4 := [4]bool{2: true, 1: true}
	arr5 := [...]bool{false, true, true, false}
	arr6 := [...]bool{3: false, 1: true, true}
	printBoolSliceType("arr1", arr1)
	printBoolSliceType("arr2", arr2)
	printBoolSliceType("arr3", arr3)
	printBoolSliceType("arr4", arr4)
	printBoolSliceType("arr5", arr5)
	printBoolSliceType("arr6", arr6)
}

func GetPointerFromContainerType() {
	println("\nGet addresses form map, slica and array::")

	pm := &map[string]int{"1": 1, "2": 2}
	ps := &[]byte{2, 56, 255}
	pa := &[...]bool{true, false, false}
	fmt.Printf("%T\n", pm)
	fmt.Printf("%T\n", ps)
	fmt.Printf("%T\n", pa)
}

func SimplifyCompositeLiterals() {
	println("\nSimplify Composite Literals::")
	heads := []*[3]byte{
		&[3]byte{'A', 'B', 'C'},
	}
	fmt.Printf("heads : %v, type %v\n", heads, reflect.TypeOf(heads))

	type LangCategory struct {
		dynamic bool
		strong  bool
	}
	langMap := map[LangCategory]map[string]int{
		LangCategory{true, true}: map[string]int{
			"Python": 1991,
			"Erlang": 1986,
		},
	}
	fmt.Printf("langMap : %v, type %v\n", langMap, reflect.TypeOf(langMap))
}

func RetrieveAndModifyContainerTypes() {
	println("\nRetrieve and Modify Container Types::")
	a := [...]int{1, 2, 3}
	s := []int{4, 5, 6}
	m := map[string]int{"A": 65, "B": 66}

	printFunc := func(a, s, m int) {
		fmt.Printf("arr[2]: %d, slice[1]: %d, map[\"B\"]: %d\n", a, s, m)
	}

	printFunc(a[2], s[1], m["B"])
	a[2], s[1], m["B"] = 9, 10, 90
	printFunc(a[2], s[1], m["B"])

	v, f := m["A"]
	fmt.Printf("m[\"A\"] = [ k: %v, key in map: %v ]\n", v, f)
	m = nil
	fmt.Println(m["A"]) // 0

	// Each of the following lines can cause a panic.
	// _ = a[v]         // panic: index out of range
	// _ = s[v]         // panic: index out of range
	// m["hello"] = 555 // panic: assign to entry in nil map
}

func ContainerAssignments() {
	println("\nContainer Assignments::")
	m0 := map[int]int{1: 12, 2: 67, 3: 99}
	m1 := m0
	m1[2] = 21
	fmt.Println("init m0: 1:12, 2:67, 3:99", "m0: ", m0, " m1: ", m1)

	s0 := []rune{'a', 'e', 'z'}
	s1 := s0
	s1[len(s1)-1] = 'w'
	fmt.Printf("init s0: 'a', 'e', 'z' :: s0: %c, s1: %c\n", s0, s1)

	a0 := [...]int{7, 8, 9}
	a1 := a0
	a1[0] = 2
	fmt.Println("a0: ", a0, " a1: ", a1) // [7 8 9] [2 8 9]
}

func AppendAndDeleteFromMap() {
	println("\nAppend and delete from Map::")
	m := map[string]int{"a": 1223, "f": 443, "j": 0}
	jm, _ := json.Marshal(m)
	fmt.Printf("Map m before modifing: %s\n", jm)
	m["j"] = 1
	jm, _ = json.Marshal(m)
	fmt.Printf("Fix entry j:0 : %s\n", jm)

	delete(m, "j")
	jm, _ = json.Marshal(m)
	fmt.Printf("Delete entry j:1 : %s\n", jm)
}

func SliceAppend() {
	println("\nSlice append::")

	printSlice := func(name string, s []int) {
		fmt.Printf("%s: %v, len: %d, cap %d\n", name, s, len(s), cap(s))
	}

	s := []int{2, 3, 4}
	// d := []int{1, 2, 3}
	_ = append(s[:0], []int{1, 2, 3}...)
	printSlice("s", s)

	s0 := []int{16, 34, 56}
	printSlice("s0", s0)

	s1 := append(s0, 87, 90)
	printSlice("s1", s1)

	s3 := append(s1, s1...)
	printSlice("s1", s3)

	s4 := append(s0) // s4 = s0
	printSlice("s4", s4)

	s5 := append(s0[:1], 111)

	s0[0] = 17
	s1[0] = 18
	s5[1] = 112

	printSlice("s0 mod", s0)
	printSlice("s1 mod", s1)
	printSlice("s3 mod", s3)
	printSlice("s4 mod", s4)
	printSlice("s5", s5)
}

func MakeSliceAndMap() {
	println("\nMake Slice and Map::")
	m := make(map[string]int)
	fmt.Println("m := make(map[string]int)", m, " len: ", len(m))

	m1 := make(map[string]int, 3)
	m1["key_1"] = 1
	m1["key_2"] = 2
	fmt.Println("m1 := make(map[string]int, 3)", m1, " len: ", len(m1))

	s := make([]int, 3)
	fmt.Println("s := make([]int)", s, " len: ", len(s), " cap: ", cap(s))

	s1 := make([]int, 3, 5)
	fmt.Println("s1 := make([]int, 3, 5)", s1, " len: ", len(s1), " cap: ", cap(s1))
}

func Reslicing() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6}
	s0 := a[:]    // a[0:7:7] a[0:len(a)], len(a) = cap(a)
	s1 := s0[:]   // s1 := s0 ~ s1 := append(s0)
	s2 := s1[1:3] // a[1:3]
	s3 := s1[3:]  // s1[3:7] ~ s1[3:] ~ a[3:cap(a)]
	s4 := s0[4:5] // s0[4:5:7] ~ a[4:5]
	s41 := s0[4:5:7]
	s42 := a[4:5]
	s5 := s0[:2:2] // s0[0:2:2] ~ a[:2] ~ a[0:2]
	println(s2, s3, s4, s41, s42, s5)
}

func CopySlices() {
	println("\nCopy slices::")
	type Ta []int
	type Tb []int
	dest := Ta{1, 2, 3}
	src := Tb{5, 6, 7, 8, 9}
	fmt.Printf("dest: %v, src: %v\n", dest, src)

	count := copy(dest, src)
	fmt.Printf("copy(dest, src) :: count: %d, dest slice: %v\n", count, dest)

	a := [4]int{}
	count = copy(a[:], src)
	fmt.Printf("copy(a[:], src) :: count: %d, dest slice: %v\n", count, a)

	count = copy(a[:], a[2:])
	fmt.Printf("copy(a[:], a[2:]) :: count: %d, dest slice: %v\n", count, a)

	b := Tb{10, 12, 23}
	count = Copy(b, src)
	fmt.Printf("Copy(b, src) :: count: %d, dest slice: %v\n", count, b)

}

// if len(dest) equals 0, then no elements will be copied
func Copy(dest, src []int) int {
	if len(dest) < len(src) {
		_ = append(dest[:0], []int{1, 2, 3}...)
		return len(dest)
	} else {
		_ = append(dest[:0], src...)
		return len(dest)
	}
}

func RangeOverArrayModificationHasNoEffects() {
	println("\nRange over array modification has no effects::")
	type Person struct {
		name string
		age  byte
	}

	persons := [2]Person{{"Ivan", 24}, {"Oleg", 25}}
	for i, v := range persons {
		fmt.Println("Before mod elem: [ ", i, " : ", v, " ]")

		// This modification has no effects on
		// the iteration, for the ranged array
		// is a copy of the persons array.
		persons[1].name = "Jack"
		fmt.Println("Fix persons[1].name = \"Jack\"")

		// This modification has not effects on
		// the persons array, for p is just a
		// copy of a copy of one persons element.
		v.age = 31

		fmt.Println("After  mod elem: [ ", i, " : ", v, " ]")
	}
	fmt.Println("Original persons array:", &persons)
	fmt.Println("\nLesson: ", "The copy of the array doesn't share elements with the array,", "\n\t so the modifications made on the array elements during the iteration will not be reflected to the iteration variables")
}

func RangeOverSliceModificationHasEffects() {
	println("\nRange over slice modification has effects::")
	type Person struct {
		name string
		age  byte
	}

	persons := []Person{{"Ivan", 24}, {"Oleg", 25}}
	for i, v := range persons {
		fmt.Println("Before mod elem: [ ", i, " : ", v, " ]")

		// Now this modification has effects
		// on the iteration.
		persons[1].name = "Jack"
		fmt.Println("Fix persons[1].name = \"Jack\"")

		// This modification still has not
		// any real effects.
		v.age = 31

		fmt.Println("After  mod elem: [ ", i, " : ", v, " ]")
	}
	fmt.Println("Original persons array:", &persons)
	fmt.Println("\nLesson: ", "The modification on the slice during the iteration has effects on the iteration,", "\n\t but the modification on the iteration variable still has no effects on the slice")
}

func RangeOverMap() {
	println("\nRange over map::")
	langs := map[struct{ dynamic, strong bool }]map[string]int{
		{true, false}:  {"JavaScript": 1995},
		{false, true}:  {"Go": 2009},
		{false, false}: {"C": 1972},
	}
	// The key type and element type of this map
	// are both pointer types. Some weird, just
	// for education purpose.
	// Если тип key и elem не pointer type, то не будет эффекта обращения к одной и той же паре key-elem
	// То есть нету обращения к одному и тому же сегменту памяти для мапы
	m0 := map[*struct{ dynamic, strong bool }]*map[string]int{}
	for category, langInfo := range langs {
		m0[&category] = &langInfo
		// This following line has no effects on langs.
		category.dynamic, category.strong = true, true
	}
	fmt.Println("Map langs:: ")
	for category, langInfo := range langs {
		fmt.Println(category, langInfo)
	}
	fmt.Println("\nMap m0:: ")
	for category, langInfo := range m0 {
		fmt.Println(category, langInfo)
	}

	m1 := map[struct{ dynamic, strong bool }]map[string]int{}
	for category, langInfo := range m0 {
		m1[*category] = *langInfo
	}
	// m0 and m1 both contain only one entry.
	fmt.Println(len(m0), len(m1)) // 1 1
	fmt.Println(m1)               // map[{true true}:map[C:1972]]
}

func UseArrayPointer() {
	println("\nUse Array Pointer Type to range over::")
	var arr [4]int = [4]int{1, 2, 3, 4}

	println("\nTake pointer:: &arr")
	for i, n := range &arr {
		fmt.Println(i, n)
	}

	println("\nTake slice:: arr[:]")
	for i, n := range arr[:] {
		fmt.Println(i, n)
	}
}

func OverArrayWithPanic() {
	println("\nOver array with panic::")
	var p *[5]int // nil

	for i, _ := range p { // okay
		fmt.Println(i)
	}

	for i := range p { // okay
		fmt.Println(i)
	}

	// for i, n := range p { // panic
	// 	fmt.Println(i, n)
	// }
}

func IndexArrayElementWithPointer() {
	println("\nIndex array element with pointer::")
	a := [...]int{1, 2, 3, 45, 67}
	p := &a
	p[3], p[4] = 4, 5
	fmt.Println(*p)
}

func DeriveSlicesFromArrayPointer() {
	println("\nDerive slices from array pointer::")
	a := &[5]int{1, 2, 3, 4, 5}
	sa := a[:3]
	fmt.Println(sa)
}

// Assume T is a small-size type.
func DeleteElements(s []int, keep func(int) bool, clear bool) []int {
	//result := make([]T, 0, len(s))
	result := s[:0] // without allocating a new slice
	for _, v := range s {
		if keep(v) {
			result = append(result, v)
		}
	}
	if clear { // avoid memory leaking
		temp := s[len(result):]
		for i := range temp {
			// t0 is a zero value literal of T.
			temp[i] = 0
		}
	}
	return result
}

func StackOper() {
	s := []int{1, 2, 3, 4}
	e := -1

	// Pop front (shift):
	s, e = s[1:], s[0]
	// Pop back:
	s, e = s[:len(s)-1], s[len(s)-1]
	// Push front
	s = append([]int{e}, s...)
	// Push back:
	s = append(s, e)

}
