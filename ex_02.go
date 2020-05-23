package main

import "fmt"

func main2() {
	const c int = 2
	// 덧셈(+)
	a := 1 + 2               // 두 정수 더하면서 변수 선언
	b := a + c               // 두 변수 더하면서 변수 선언
	str := "Hello, " + "Go!" // 두 문자열 붙이면서 새 문자열 선언

	// 뺄셈(-)
	d := a - b // 두 변수 빼면서 변수 선언

	// 곱셈(*), 나눗셈(/), 나머지(%)
	e := a * b // 두 변수 곱하면서 변수 선언
	f := e / a // 두 변수 나누면서 변수 선언
	g := b % 2 // 나머지 연산 하며 변수 선언

	fmt.Println(a, b, c, d, e, f, g, str)
}
