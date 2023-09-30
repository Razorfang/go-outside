package main

import (
	"fmt"
)

type ShopItem struct {
	name string
	priceCents int
}

type Purchasable interface {
	explain()
}

type Fruit struct {
	ShopItem
	unitWeightMilligrams int
}

func newFruit(name string, priceCents int, unitWeightMilligrams int) Fruit {
	return Fruit{ShopItem: ShopItem{name, priceCents}, unitWeightMilligrams: unitWeightMilligrams}
}

func (fruit Fruit) explain() {
	fmt.Printf("This fruit's name is '%s', it weight %d milligrams, and it costs %d cents.\n",
		fruit.ShopItem.name, fruit.unitWeightMilligrams, fruit.ShopItem.priceCents)
}

type Service struct {
	ShopItem
}

func newService(name string, priceCents int) Service {
	return Service{ShopItem: ShopItem{name, priceCents}}
}

func (service Service) explain() {
	fmt.Printf("This service is named '%s' and it will cost you %d cents.\n",
		service.ShopItem.name, service.ShopItem.priceCents)
}

func morshu(item Purchasable) {
	item.explain()
}

func main() {
	apple := newFruit("apple", 100, 50)
	morshu(apple)

	wrapGift := newService("wrap gift", 200)
	morshu(wrapGift)
}
