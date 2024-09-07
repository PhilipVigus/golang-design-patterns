package prototype

import (
	"errors"
	"fmt"
)

// ItemInfoGetter is an interface that includes the GetInfo method
type ItemInfoGetter interface {
	GetInfo() string
}

// ShirtCloner is an interface that includes the GetClone method
type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

// const defining the different shirt types
const (
	White = iota
	Black
	Blue
)

// GetShirtsCloner returns a new ShirtsCache
func GetShirtsCloner() ShirtCloner {
	return new(ShirtsCache)
}

// ShirtsCache is an empty structure
type ShirtsCache struct{}

// GetClone is a method defined on the ShirtsCache type
// It dereferences and returns pointers to the different prototypes, which effectively
// clones them rather than returning the prototypes themselves
func (s *ShirtsCache) GetClone(m int) (ItemInfoGetter, error) {
	switch m {
	case White:
		// clones the prototype as part of dereferencing it
		newItem := *whitePrototype
		// returns a pointer to the cloned prototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("shirt model not recognised")
	}
}

// ShirtColor is a type alias for a byte type
type ShirtColor byte

// Shirt is a struct defining what attributes a shirt has
type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

// GetInfo is an implementation of the ItemInfoGetter interface
func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("Shirt with SKU '%s' and Color id %d that costs %f\n", s.SKU, s.Color, s.Price)
}

// whitePrototype is a white shirt prototype
var whitePrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}

// blackPrototype is a black shirt prototype
var blackPrototype *Shirt = &Shirt{
	Price: 16.00,
	SKU:   "empty",
	Color: Black,
}

// bluePrototype is a blue shirt prototype
var bluePrototype *Shirt = &Shirt{
	Price: 17.00,
	SKU:   "empty",
	Color: Blue,
}

// GetPrice returns the price of a shirt
func (s *Shirt) GetPrice() float32 {
	return s.Price
}
