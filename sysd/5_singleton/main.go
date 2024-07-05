package main

import "sync"

func main() {
  boiler := NewChocolateBoiler()

  var wg sync.WaitGroup
  wg.Add(3)

  go func ()  {
    defer wg.Done()
    boiler.Fill()
  }()

  go func ()  {
    defer wg.Done()
    boiler.Boil()
  }()

  go func() {
    defer wg.Done()
    boiler.Drain()
  }()

  wg.Wait()
}
