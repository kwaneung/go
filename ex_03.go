package main

import "fmt"

func main3() {
	var a int = 0x1020
	var b int = 0x1111
	shift := b << 1
	shift2 := a >> 1
	and := a & b
	or := a | shift
	xor := a ^ b
	fmt.Printf("a = %d(0x%x(%b))\nb = %d(0x%x(%b))\n", a, a, a, b, b, b)
	fmt.Printf("shift = %x (%b,%d)\n", shift, shift, shift)
	fmt.Printf("shift2 = %x (%b,%d)\n", shift2, shift2, shift2)
	fmt.Printf("a & b = 0x%x,%b\n", and, and)
	fmt.Printf("a | b = 0x%x,%b\n", or, or)
	fmt.Printf("a ^ b = 0x%x,%b\n", xor, xor)

}
