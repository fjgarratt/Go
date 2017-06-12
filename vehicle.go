package main

import (
    "fmt"
)

type Vehicle struct{
        Started bool
        Consume float32
        Speed   float32
}
type Volta struct {
        Vehicle
}
type Propeller interface {
        ConsumeEnergy() float32
        GetSpeed()      float32
}

func (v *Vehicle) ConsumeEnergy() float32 {
    v.Consume += 30
    v.Speed   += 15
    fmt.Println("Speed is now", v.Speed)
    return v.Consume
}

func (v *Vehicle) GetSpeed() float32 {

    return v.Speed
}

func (v *Vehicle) Start() {
    v.Started = true
    fmt.Println("Vehicle started:", v.Started)
}

func (v Volta) Start() {
    v.Vehicle.Start()
    fmt.Println("Volta started")
}

func main() {
    veh := &Vehicle{}
    veh.Start()
    
    var prp Propeller = veh
    
    prp.ConsumeEnergy()
    prp.ConsumeEnergy()
    
    consumed := prp.ConsumeEnergy()
    
    fmt.Println("Energy Consumed is ", consumed )
    fmt.Println("Speed is ", prp.GetSpeed() )
    
}