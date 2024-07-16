package main

import "fmt"

type Beverage interface {
  GetDescription() string
  Cost() float64
}

type BaseBeverage struct {
  description string
}

func (b *BaseBeverage) GetDescription() string {
  return b.description
}

func (b *BaseBeverage) Cost() float64 {
  return 0.0
}

func main() {
  espresso := NewEspresso()
  fmt.Printf("%s: $%.2f\n", espresso.GetDescription(), espresso.Cost())

  houseblend := NewHouseBlend()
  houseblend = NewMocha(houseblend)
  houseblend = NewWhip(houseblend)
  fmt.Printf("%s: $%.2f\n", houseblend.GetDescription(), houseblend.Cost())
}
