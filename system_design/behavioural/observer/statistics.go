package main

import (
	"fmt"
	"math"
)

type StatisticsDisplay struct {
  avgTemp float64
  minTemp float64
  maxTemp float64
  temperatures []float64
  weatherData Subject
}

func NewStatisticsDisplay(wd Subject) *StatisticsDisplay {
  display := &StatisticsDisplay{
    weatherData: wd,
    minTemp: math.Inf(1),
    maxTemp: math.Inf(-1),
  }
  wd.RegisterObserver(display)
  return display
}

func (sd *StatisticsDisplay) Update(temp, humid, pres float64) {
  sd.temperatures = append(sd.temperatures, temp)
  sd.calculateStatistics()
  sd.Display()
}

func (sd *StatisticsDisplay) Display() {
  fmt.Printf("Avg/Max/Min temperature = %.2f/%.2f/%.2f\n", sd.avgTemp, sd.maxTemp, sd.minTemp)
}

func (sd *StatisticsDisplay) calculateStatistics() {
  sum := 0.0
  for _, temp := range sd.temperatures {
    sum += temp
    if temp < sd.minTemp {
      sd.minTemp = temp
    }
    if temp > sd.maxTemp {
      sd.maxTemp = temp
    }
  }

  sd.avgTemp = sum / float64(len(sd.temperatures))
}
