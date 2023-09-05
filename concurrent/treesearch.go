package main

import (
	"fmt"
)

type Fruit struct {
	name string
	description string
	weight float32
}

func (f Fruit) String() string {
	return fmt.Sprintf("The name of this fruit is '%s'. It is %s", f.name, f.description)
}

type Tree[T any] struct {
	left *Tree[T]
	value T
	right *Tree[T]
}

func (tree *Tree[T]) insert(value T) {
	fmt.Println(value)
}

func main() {
	apple := Fruit{"apple", "a round, crisp specimen with varying sweetness.", 100.0}
	orange := Fruit{"orange", "somewhat more spherical than the apple.", 95.0}
	mangosteen := Fruit{"mangosteen", "really something I guess.", 30.0}

	mangosteenNode := Tree[Fruit]{nil, mangosteen, nil}
	orangeNode := Tree[Fruit]{nil, orange, nil}
	fruitTree := Tree[Fruit]{&mangosteenNode, apple, &orangeNode}

	fmt.Println(fruitTree.value)
	fruitTree.insert(Fruit{"blueberry", "a minuscule, tart little thing.", 0.01})
}
