package main

import "fmt"

const (
	adidas = "adidas"
	nike   = "nike"
)

type SportsFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func GetSportsFactory(brand string) (SportsFactory, error) {
	if brand == adidas {
		return &Adidas{}, nil
	}

	if brand == nike {
		return &Nike{}, nil
	}

	return nil, fmt.Errorf("wrong brand type passed")
}
