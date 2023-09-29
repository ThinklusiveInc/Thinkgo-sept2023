package main

import (
	"fmt"

	"golang.org/x/exp/maps"
)

func main() {
	m := make(map[string]int)
	m["h1"] = 230
	m["h2"] = 240
	fmt.Println("map:", m)
	v1 := m["h1"]
	fmt.Println("v1:", v1)
	fmt.Println(len(m))
	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
