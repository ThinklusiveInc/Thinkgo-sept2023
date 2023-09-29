package main

import (
	"fmt"
)

func main() {
	//var s []string
	//fmt.Println("uninit", s, s == nil, len(s) == 0)

	s := make([]string, 4)
	fmt.Println(s, len(s), cap(s))
	s[0] = "h"
	s[1] = "q"
	s[2] = "b"
	s[3] = "C"
	fmt.Println("set", s)
	fmt.Println("get", s[2])

	s = append(s, "s")
	s = append(s, "g", "k")
	fmt.Println("append:", s)

	c := make([]string, 5)
	copy(c, s)
	fmt.Println("copy c:", c)
	j := s[:5]
	fmt.Println("j1:", j)
	j = s[2:]
	fmt.Println("j2:", j)

	u := []string{"r", "e", "w", "q"}
	fmt.Println(u)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerlen := i + 1
		twoD[i] = make([]int, innerlen)
		for j := 0; j < innerlen; j++ {
			twoD[i][j] = i + j

		}
	}
	fmt.Println(twoD)

}
