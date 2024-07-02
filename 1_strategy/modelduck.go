package main

import "fmt"

type ModelDuck struct {
  Duck
}

func (m *ModelDuck) display() {
  fmt.Println("I am a model duck")
}

func NewModelDuck() *ModelDuck {
  return &ModelDuck{
    Duck: Duck{
      flyBehaviour: FlyNoWay{},
      quackBehaviour: Squeak{},
    },
  }
}

func callModelDuck() {
  modelDuck := NewModelDuck()
  modelDuck.display()
  modelDuck.performFly()
  modelDuck.performQuack()

  modelDuck.setFlyBehaviour(FlyRocketPowered{})
  modelDuck.performFly()
}
