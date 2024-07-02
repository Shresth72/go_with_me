package main

import "fmt"

// QuackBehaviour interface
type QuackBehaviour interface {
  quack()
}

// Quack struct 
type Quack struct {}

func (q Quack) quack() {
  fmt.Println("Say quack")
}

// Quack struct 
type MuteQuack struct {}

func (q MuteQuack) quack() {
  fmt.Println("<< Silence >>")
}

// Squeak struct 
type Squeak struct {}

func (q Squeak) quack() {
  fmt.Println("Say squeak")
}
