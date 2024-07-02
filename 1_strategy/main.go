package main

import "fmt"

type Duck struct {
  flyBehaviour FlyBehaviour
  quackBehaviour QuackBehaviour
}

func (d *Duck) display() {}

func (d *Duck) performFly() {
  d.flyBehaviour.fly()
}

func (d *Duck) performQuack() {
  d.quackBehaviour.quack()
}

func (d *Duck) swim() {
  fmt.Println("All ducks float")
}

func (d *Duck) setFlyBehaviour(fb FlyBehaviour) {
  d.flyBehaviour = fb
}

func (d *Duck) setQuackBehaviour(qb QuackBehaviour) {
  d.quackBehaviour = qb
}

func main() {
  // Bird with FlyWithWings and Quack Behaviour 
  mallard := Duck{
    flyBehaviour: FlyWithWings{},
    quackBehaviour: Quack{},
  }

  mallard.performFly()
  mallard.performQuack()

  // Change behaviour at runtime
  mallard.setFlyBehaviour(FlyNoWay{})
  mallard.performFly()
  fmt.Println("")

  // Model Duck
  callModelDuck()
}
