package main

import "fmt"

// Implements Observer and DisplayElement
type CurrentConditionsDisplay struct {
  temperature float64
  humidity float64
  weatherData Subject
}

func NewCurrentConditionsDisplay(wd Subject) *CurrentConditionsDisplay {
  display := &CurrentConditionsDisplay{
    weatherData: wd,
  }
  wd.RegisterObserver(display)
  return display
}

func (cd *CurrentConditionsDisplay) Update(temp, humid, pres float64) {
  cd.temperature = temp
  cd.humidity = humid
  cd.Display()
}

func (cd *CurrentConditionsDisplay) Display() {
  fmt.Printf("Current conditions: %.2f°F temperature and %.2f%% humidity\n", cd.temperature, cd.humidity)
}
