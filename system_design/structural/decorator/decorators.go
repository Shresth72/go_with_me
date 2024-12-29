package main

// Base Decorator
type CondimentDecorator struct {
  Beverage
}

// Mocha
type Mocha struct {
  beverage Beverage
}

func NewMocha(b Beverage) Beverage {
  return &Mocha{
    beverage: b,
  }
}

func (m *Mocha) GetDescription() string {
  return m.beverage.GetDescription() + ", mocha"
}

func (m *Mocha) Cost() float64 {
  return 0.20 + m.beverage.Cost()
}

// Whip
type Whip struct {
	beverage Beverage
}

func NewWhip(beverage Beverage) Beverage {
	return &Whip{
		beverage: beverage,
	}
}

func (w *Whip) GetDescription() string {
	return w.beverage.GetDescription() + ", whip"
}

func (w *Whip) Cost() float64 {
	return 0.10 + w.beverage.Cost()
}
