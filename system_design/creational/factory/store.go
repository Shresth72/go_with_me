package main

// PizzaStore interface
type PizzaStore interface {
	OrderPizza(typePizza string) Pizza
	CreatePizza(typePizza string) Pizza
}

// NYPizzaStore struct
type NYPizzaStore struct{}

func (n *NYPizzaStore) OrderPizza(typePizza string) Pizza {
	pizza := n.CreatePizza(typePizza)
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}

func (n *NYPizzaStore) CreatePizza(typePizza string) Pizza {
	if typePizza == "cheese" {
		return NewNYStyleCheesePizza()
	}
  // can add more types
	return nil
}

// ChicagoPizzaStore struct
type ChicagoPizzaStore struct{}

func (c *ChicagoPizzaStore) OrderPizza(typePizza string) Pizza {
	pizza := c.CreatePizza(typePizza)
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}

func (c *ChicagoPizzaStore) CreatePizza(typePizza string) Pizza {
	if typePizza == "cheese" {
		return NewChicagoStyleCheesePizza()
	}
	return nil
}
