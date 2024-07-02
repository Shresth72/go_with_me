package main

type WeatherData struct {
  observers []Observer
  temperature float64
  humidity float64
  pressure float64
}

func NewWeatherData() *WeatherData {
  return &WeatherData{
    observers: make([]Observer, 0),
  }
}

// Implementing Subject
func (wd *WeatherData) RegisterObserver(o Observer) {
  wd.observers = append(wd.observers, o)
}

func (wd *WeatherData) RemoveObserver(o Observer) {
  if wd.observers == nil || o == nil {
    return
  }

  for i, observer := range wd.observers {
    if observer == o {
      copy(wd.observers[i:], wd.observers[i+1:])
      wd.observers = wd.observers[:len(wd.observers)-1]
      break
    }
  }
}

func (wd *WeatherData) NotifyObservers() {
  for _, observer := range wd.observers {
    observer.Update(wd.temperature, wd.humidity, wd.pressure)
  }
}

func (wd *WeatherData) SetMeasurements(temp, humid, pres float64) {
  wd.temperature = temp
  wd.humidity = humid
  wd.pressure = pres
  wd.measurementsChanged()
}

func (wd *WeatherData) measurementsChanged() {
  wd.NotifyObservers()
}
