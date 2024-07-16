package main

import "fmt"

type Pizza interface {
  Prepare()
  Bake()
  Cut()
  Box()
  GetName() string
}

type BasePizza struct {
  Name string
  Dough string
  Sauce string
  Toppings []string
}

func (p *BasePizza) Prepare() {
  fmt.Println("Preparing", p.Name)
	fmt.Println("Tossing dough...")
	fmt.Println("Adding sauce...")
	fmt.Println("Adding toppings:")
	for _, topping := range p.Toppings {
		fmt.Println("   ", topping)
	}
}

func (p *BasePizza) Bake() {
	fmt.Println("Bake for 25 minutes at 350")
}

func (p *BasePizza) Cut() {
	fmt.Println("Cutting the pizza into diagonal slices")
}

func (p *BasePizza) Box() {
	fmt.Println("Place pizza in official PizzaStore box")
}

func (p *BasePizza) GetName() string {
	return p.Name
}

func main() {
	nyStore := &NYPizzaStore{}
	chicagoStore := &ChicagoPizzaStore{}

	pizza := nyStore.OrderPizza("cheese")
	fmt.Println("Ethan ordered a", pizza.GetName())

	pizza = chicagoStore.OrderPizza("cheese")
	fmt.Println("Joel ordered a", pizza.GetName())
}
