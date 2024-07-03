package main

import "fmt"

// FlyBehaviour interface
type FlyBehaviour interface {
  fly()
}

// FlyWithWings struct 
type FlyWithWings struct {}

func (f FlyWithWings) fly() {
  fmt.Println("I can fly")
}

// FlyNoWay struct 
type FlyNoWay struct {}

func (f FlyNoWay) fly() {
  fmt.Println("I can't fly")
}

// FlyRocketPowered struct 
type FlyRocketPowered struct {}

func (f FlyRocketPowered) fly() {
  fmt.Println("I can fly with rockets")
}
