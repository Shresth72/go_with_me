package main

// Espresso
type Espresso struct {
  BaseBeverage
}

func NewEspresso() Beverage {
  return &Espresso{
    BaseBeverage{
      description: "espresso",
    },
  }
}

func (e *Espresso) Cost() float64 {
  return 1.99
}

// HouseBlend
type HouseBlend struct {
  BaseBeverage
}

func NewHouseBlend() Beverage {
  return &HouseBlend{
    BaseBeverage{
      description: "house blend coffee",
    },
  }
}

func (e *HouseBlend) Cost() float64 {
  return 0.89
}
