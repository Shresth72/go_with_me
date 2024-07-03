package main

import "fmt"

// NYStyle
type NYStyleCheesePizza struct {
  BasePizza
}

func NewNYStyleCheesePizza() Pizza {
  return &NYStyleCheesePizza{
		BasePizza: BasePizza{
			Name:     "NY Style Sauce and Cheese Pizza",
			Dough:    "Thin Crust Dough",
			Sauce:    "Marinara Sauce",
			Toppings: []string{"Grated Reggiano Cheese"},
		},
	}
}

// ChicagoStyleCheesePizza struct
type ChicagoStyleCheesePizza struct {
	BasePizza
}

func NewChicagoStyleCheesePizza() Pizza {
	return &ChicagoStyleCheesePizza{
		BasePizza: BasePizza{
			Name:     "Chicago Style Deep Dish Cheese Pizza",
			Dough:    "Extra Thick Crust Dough",
			Sauce:    "Plum Tomato Sauce",
			Toppings: []string{"Shredded Mozzarella Cheese"},
		},
	}
}

func (p *ChicagoStyleCheesePizza) Cut() {
	fmt.Println("Cutting the pizza into square slices")
}


