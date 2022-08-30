// GO 101 articles
// basic types: https://go101.org/article/basic-types-and-value-literals.html

package types

import (
	"fmt"
	"math/bits"
	"reflect"
	"unsafe"
)

func IntSizeBits() {
	//This is computed as const uintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64
	sizeOfIntInBits := bits.UintSize
	fmt.Printf("SizeOfIntInBits: %d bits\n", sizeOfIntInBits)

	var a int
	fmt.Printf("%d bytes\n", unsafe.Sizeof(a))
	fmt.Printf("a's type is %s\n", reflect.TypeOf(a))
}

func FpValues() {
	var (
		f1 float32 = 1.23
		f2 float32 = 01.23
		f3 float32 = .23

		fe1  float32 = 1.23e2  // == 123.0
		fe2  float32 = 123e2   // == 12300.0
		fe3  float32 = 123e2   // == 12300.0
		fe4  float32 = 123.e2  // == 12300.0
		fe5  float32 = 1e-1    // == 0.1
		fe6  float32 = .1e0    // == 0.1
		fe7  float32 = 0010e-2 // == 0.1
		fez1 float32 = 0e5     // 0.0
		fez2 float32 = 0e0
	)

	fmt.Printf("%v\n", f1)
	fmt.Printf("%v\n", f2)
	fmt.Printf("%v\n", f3)
	fmt.Printf("%v\n", fe1)
	fmt.Printf("%v\n", fe2)
	fmt.Printf("%v\n", fe3)
	fmt.Printf("%v\n", fe4)
	fmt.Printf("%v\n", fe5)
	fmt.Printf("%v\n", fe6)
	fmt.Printf("%v\n", fe7)
	fmt.Printf("%v\n", fez1)
	fmt.Printf("%v\n", fez2)
}

func FpHexValues() {
	var (
		fp1 = 0x1p-2     // == 1.0/4 = 0.25
		fp2 = 0x2.p10    // == 2.0 * 1024 == 2048.0
		fp3 = 0x1.Fp+0   // == 1+15.0/16 == 1.9375
		fp4 = 0x.8p1     // == 8.0/16 * 2 == 1.0
		fp5 = 0x1FFFp-16 // == 0.1249847412109375
		fpz = 0x0p0
	)

	/*
		illegal statements

		0x.p1    // mantissa has no digits
		1p-2     // p exponent requires hexadecimal mantissa
		0x1.5e-2 // hexadecimal mantissa requires p exponent
	*/

	fmt.Printf("%v\n", fp1)
	fmt.Printf("%v\n", fp2)
	fmt.Printf("%v\n", fp3)
	fmt.Printf("%v\n", fp4)
	fmt.Printf("%v\n", fp5)
	fmt.Printf("%v\n", fpz)
}

func RuneValues() {
	println('a' == 97)
	println('a' == '\141') // 141 is the octal representation of decimal number 97.
	println('a' == '\x61') // 61 is the hex representation of decimal number 97.
	println('a' == '\u0061')
	println('a' == '\U00000061')
	println(0x61 == '\x61')
	println('\u4f17' == '众')

	var simpleRune rune = '£'
	fmt.Printf("Rune value - type: %s\n", reflect.TypeOf(simpleRune))
	fmt.Printf("Unicode CodePoint: %U\n", simpleRune)
	fmt.Printf("Character: %c\n", simpleRune)

	s := "0b£"
	fmt.Printf("\nString %s as rune\n", s)

	//This will print the Unicode Points
	fmt.Printf("%U\n", []rune(s))

	//This will the decimal value of Unicode Code Point
	fmt.Printf("%d\n", []rune(s))
}

func StringValues() {
	// The following interpreted string literals are equivalent.
	s1 := "\u4f17\xe4\xba\xba"
	// The Unicode of 众 is 4f17, which is
	// UTF-8 encoded as three bytes: e4 bc 97.
	s2 := "\xe4\xbc\x97\u4eba"
	// The Unicode of 人 is 4eba, which is
	// UTF-8 encoded as three bytes: e4 ba ba.
	s3 := "\xe4\xbc\x97\xe4\xba\xba"

	fmt.Printf("S1 string value: %s, code points:%U, dec code points:%d\n", s1, []byte(s1), []byte(s1))
	fmt.Printf("S2 string value: %s, code points:%U, dec code points:%v\n", s2, []byte(s2), []byte(s2))
	fmt.Printf("S3 string value: %s, code points:%U, dec code points:%v\n", s3, []byte(s3), []byte(s3))

	s4 := "ab£"
	fmt.Printf("S4 string value: %s, code points: %U, dec code points: %v\n", s4, []rune(s4), []byte(s4))

	s5 := 12 + 'A'
	fmt.Printf("S5 string value: %v, type: %v\n", s5, reflect.TypeOf(s5))

}

func ConstValues() {
	const uintBitSize1 uint = (1 << 64) - 1
	fmt.Printf("UintBitSize1 ~ (1 << 64) - 1: %d bits\n", uintBitSize1)

	const uintBitSize2 uint = 32 << (^uint(0) >> 32 & 1) // 32 or 64
	fmt.Printf("UintBitSize2 ~ 32 << (^uint(0) >> 32 & 1) : bits.UintSize: %d bits\n", uintBitSize2)

	fmt.Printf("\nThe largest uint value ~ ^uint(0): %v\n", ^uint(0))
	fmt.Printf("The largest int value ~ int(^uint(0) >> 1): %v\n", int(^uint(0)>>1))

	const NativeWordBits = 32 << (^uint(0) >> 63)
	const Is64bitOS = ^uint(0)>>63 != 0
	const Is32bitOS = ^uint(0)>>32 == 0

	fmt.Printf("\nIs64bitOS ~ ^uint(0) >> 63 != 0 : %v\n", Is64bitOS)
	fmt.Printf("Is32bitOS ~ ^uint(0) >> 32 == 0 : %v\n", Is32bitOS)
	fmt.Printf("NativeWordBits ~ 32 << (^uint(0) >> 63): %v\n", NativeWordBits)

	fmt.Printf("\nConst iota::\n")

	const (
		k = 3 // now, iota == 0

		m float32 = iota + .5 // m float32 = 1 + .5
		n                     // n float32 = 2 + .5

		p    = 9          // now, iota == 3
		q    = iota * 2   // q = 4 * 2
		_                 // _ = 5 * 2
		r                 // r = 6 * 2
		s, t = iota, iota // s, t = 7, 7
		u, v              // u, v = 8, 8
		_, w              // _, w = 9, 9
	)

	const x = iota // x = 0
	const (
		y = iota // y = 0
		z        // z = 1
	)

	const (
		Failed    = iota - 1 // == -1
		Unknown              // == 0
		Succeeded            // == 1
	)

	const (
		Readable   = 1 << iota // == 1
		Writable               // == 2
		Executable             // == 4
	)

	println(k)
	println(m)             // +1.500000e+000
	println(n)             // +2.500000e+000
	println(q, r)          // 8 12
	println(s, t, u, v, w) // 7 7 8 8 9
	println(x, y, z)       // 0 0 1

	fmt.Printf("\nFailed const: %v\n", Failed)
	fmt.Printf("Unknown const: %v\n", Unknown)
	fmt.Printf("Succeeded const: %v\n", Succeeded)
	fmt.Printf("\nReadable const: %v\n", Readable)
	fmt.Printf("Writable const: %v\n", Writable)
	fmt.Printf("Executable const: %v\n", Executable)

	// 	const a = -1.23
	// 	// The type of b is deduced as float64.
	// 	var b = a
	// 	// error: constant 1.23 truncated to integer.
	// 	var x1 = int32(a)
	// 	// error: cannot assign float64 to int32.
	// 	var y1 int32 = b
	// 	// okay: z == -1, and the type of z is int32.
	// 	//       The fraction part of b is discarded.
	// 	var z1 = int32(b)

	// 	const k1 int16 = 255
	// 	// The type of n is deduced as int16.
	// 	var n1 = k1
	// 	// error: constant 256 overflows uint8.
	// 	var f1 = uint8(k1 + 1)
	// 	// error: cannot assign int16 to uint8.
	// 	var g1 uint8 = n1 + 1
	// 	// okay: h == 0, and the type of h is uint8.
	// 	//       n+1 overflows uint8 and is truncated.
	// 	var h1 = uint16(n1 + 1)
}

// 3 untyped named constants. Their bound
// values all overflow their respective
// default types. This is allowed.
const utn = 1 << 64          // overflows int
const utr = 'a' + 0x7FFFFFFF // overflows rune
const utx = 2e+308           // overflows float64

func UntypedNamedConst() {
	_ = utn >> 2
	_ = utr - 0x7FFFFFFF
	_ = utx / 2
}
