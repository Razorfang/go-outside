package main

import (
	"fmt"
)

type rectangle struct {
	width uint64
	height uint64
}

func (r *rectangle) get_area() uint64 {
	return r.width * r.height
}

func get_area_2(r *rectangle) uint64 {
	return r.width * r.height
}

func main() {
	r  := rectangle{3, 2}
	fmt.Println(r.get_area())
	fmt.Println(get_area_2(&r))
}
