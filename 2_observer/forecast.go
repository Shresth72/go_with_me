package main

import "fmt"

type ForecastDisplay struct {
  temperature float64
  humidity float64
  weatherData Subject
}

func NewForecastDisplay(wd Subject) *ForecastDisplay {
  display := &ForecastDisplay{
    weatherData: wd,
  }
  wd.RegisterObserver(display)
  return display
}

func (fd *ForecastDisplay) Update(temp, humid, _ float64) {
  fd.temperature = temp
  fd.humidity = humid
  fd.Display()
}

func (fd *ForecastDisplay) Display() {
	forecast := fd.generateForecast()
	fmt.Println("Forecast:", forecast)
  println("\n")
}

func (fd *ForecastDisplay) generateForecast() string {
	if fd.temperature > 80 && fd.humidity > 60 {
		return "Expect hot and humid conditions."
	} else if fd.temperature < 50 {
		return "It's likely to be cold."
	} else {
		return "Weather conditions look moderate."
	}
}
