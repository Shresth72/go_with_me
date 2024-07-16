package main

import "fmt"

// Target
type mobile interface {
  chargeAppleMobile()
}

// Concrete Impl 
type apple struct {}

func (a *apple) chargeAppleMobile() {
  fmt.Println("charging apple device")
}

// Client
type client struct {}

func (a *client) ChargeMobile(mob mobile) {
  mob.chargeAppleMobile()
}

// Adaptee
type android struct {}

func (a *android) chargeAndroidMobile() {
  fmt.Println("charging android device")
}

// Adapter (convert the input into a form for anabling android)
type androidAdapter struct {
  android *android
}

func (aa *androidAdapter) chargeAppleMobile() {
  aa.android.chargeAndroidMobile()
}
