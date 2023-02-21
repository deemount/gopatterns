package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtColor byte

type Shirt struct {
	Price float32
	SN    string
	Color ShirtColor
}

func (i *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs %f", i.SN, i.Color, i.Price)
}

var whitePrototype = &Shirt{
	Price: 15.00,
	SN:    uuid.New().String(),
	Color: White,
}

var blackPrototype = &Shirt{
	Price: 16.00,
	SN:    uuid.New().String(),
	Color: Black,
}

var bluePrototype = &Shirt{
	Price: 17.00,
	SN:    uuid.New().String(),
	Color: Blue,
}

func (i *Shirt) GetPrice() float32 {
	return i.Price
}

const (
	White = 1
	Black = 2
	Blue  = 3
)

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

func GetShirtsCloner() ShirtCloner {
	return new(ShirtsColors)
}

type ShirtsColors struct{}

func (sc *ShirtsColors) GetClone(s int) (ItemInfoGetter, error) {
	switch s {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("Shirt model not recognized")
	}
}
