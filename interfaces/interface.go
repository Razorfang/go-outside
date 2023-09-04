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

func (p person) String() string {
	return fmt.Sprintf("Height: [%f] Weight: [%f] Walking Speed: [%f] Distance Travelled: [%f]", p.height, p.weight, p.walking_speed, p.distance_travelled)
}

func main() {
	jamie := person{1.2, 3.4, 5.6, 0}
	fmt.Println(jamie)
	jamie.move_forward(1)
	fmt.Println(jamie)
}
