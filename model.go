package main

const (
	FirstConst = "test"
)

var (
	FirstVar = 1
)

// start
type Vehicle struct {
	Color string
}

func NewVehicle(color string) *Vehicle {
	return &Vehicle{Color: color}
}

func (v Vehicle) GetColor() string {
	return v.Color
}

// end

func Something() {
	// something here
}
