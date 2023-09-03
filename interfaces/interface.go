package main

import (
	"fmt"
)

type person struct {
	height float32
	weight float32
	walking_speed float32
	distance_travelled float32
}

type Walker interface {
	move_forward(distance float32) float32
}

func (p *person) move_forward(distance float32) float32 {
	p.distance_travelled += p.walking_speed * distance
	return p.distance_travelled
}

func main() {
	jamie := person{1.2, 3.4, 5.6, 0}
	fmt.Println(jamie.move_forward(1))
}
