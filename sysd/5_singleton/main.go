package main

import (
	"sync"
	"time"
)

func main() {
	boiler := NewChocolateBoiler()

	for {
		var wg sync.WaitGroup
		wg.Add(3)

		go func() {
			defer wg.Done()
			err := boiler.Fill()
			if err != nil {
				println(err)
				return
			}
		}()
		time.Sleep(1 * time.Second)

		go func() {
			defer wg.Done()
			err := boiler.Boil()
			if err != nil {
				println(err)
				return
			}
		}()
		time.Sleep(1 * time.Second)

		go func() {
			defer wg.Done()
			err := boiler.Drain()
			if err != nil {
				println(err)
				return
			}
		}()
		time.Sleep(1 * time.Second)

		wg.Wait()
	}
}
