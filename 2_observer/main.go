package main

// Interfaces
type Subject interface {
  RegisterObserver(o Observer)
  RemoveObserver(o Observer)
  NotifyObservers()
}

type Observer interface {
  Update(temp float64, humidity float64, pressure float64)
}

type DisplayElement interface {
  Display()
}

func main() {
  weatherData := NewWeatherData()
  conditionDisplay := NewCurrentConditionsDisplay(weatherData)
  statisticsDisplay := NewStatisticsDisplay(weatherData)
  forecastDisplay := NewForecastDisplay(weatherData)

  // Weather changes
  weatherData.SetMeasurements(75.0, 60.0, 30.4)
  weatherData.SetMeasurements(80.0, 65.0, 30.2)
  weatherData.SetMeasurements(82.0, 70.9, 30.1)

  weatherData.RemoveObserver(conditionDisplay)
  weatherData.RegisterObserver(statisticsDisplay)
  weatherData.RegisterObserver(forecastDisplay)
}
