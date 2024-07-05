package main

import (
	"fmt"
	"sync"
)

type Boiler interface {
	Fill()
	Drain()
	Boil()
	IsEmpty() bool
	IsBoiled() bool
}

type ChocolateBoiler struct {
	empty  bool
	boiled bool
	mu     sync.Mutex
	cond   *sync.Cond
}

var once sync.Once
var instance *ChocolateBoiler

// Returns the singleton instance
func NewChocolateBoiler() *ChocolateBoiler {
	once.Do(func() {
		instance = &ChocolateBoiler{
			empty:  true,
			boiled: false,
		}
		instance.cond = sync.NewCond(&instance.mu)
	})
	return instance
}

func (c *ChocolateBoiler) Fill() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for !c.IsEmpty() {
		c.cond.Wait()
	}

	c.empty = false
	c.boiled = false

	// full the boiler with milk/choco mix
	fmt.Println("Boiler is filled")
	c.cond.Broadcast()
}

func (c *ChocolateBoiler) Drain() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for c.IsEmpty() || !c.IsBoiled() {
		c.cond.Wait()
	}

	// drain the boiled contents
	c.empty = true
	fmt.Println("Boiler is drained")
	c.cond.Broadcast()
}

func (c *ChocolateBoiler) Boil() {
	c.mu.Lock()
	defer c.mu.Unlock()

	for c.IsEmpty() || c.IsBoiled() {
		c.cond.Wait()
	}

	// bring the contents to boil
	c.boiled = true
	fmt.Println("Boiler is boiled")
	c.cond.Broadcast()
}

func (c *ChocolateBoiler) IsEmpty() bool {
	return c.empty
}

func (c *ChocolateBoiler) IsBoiled() bool {
	return c.boiled
}
