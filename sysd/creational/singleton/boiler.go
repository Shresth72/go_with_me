package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
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

const (
	timeForRequest  = 1 * time.Second
	timeoutDuration = 1 * time.Second
	timoutError     = "timout occurred"
)

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

func (c *ChocolateBoiler) Fill() error {
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
	return nil
}

func (c *ChocolateBoiler) Boil() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	for c.IsEmpty() {
		c.cond.Wait()
	}

	// bring the contents to boil
	c.boiled = true
	fmt.Println("Boiler is boiled")
	c.cond.Broadcast()
	return nil
}

func (c *ChocolateBoiler) Drain() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	timeout := time.After(2 * time.Second)
	for c.IsEmpty() || !c.IsBoiled() {
		select {
		case <-timeout:
			return errors.New("timeout while waiting to drain")
		default:
			c.cond.Wait()
		}
	}

	// drain the boiled contents
	c.empty = true
	c.boiled = false
	fmt.Println("Boiler is drained")
	c.cond.Broadcast()
	return nil
}

// Utils
func (c *ChocolateBoiler) IsEmpty() bool {
	return c.empty
}

func (c *ChocolateBoiler) IsBoiled() bool {
	return c.boiled
}

// Setter for tests
func (c *ChocolateBoiler) SetEmpty(empty bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.empty = empty
}

func (c *ChocolateBoiler) SetBoiled(boiled bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.boiled = boiled
}
