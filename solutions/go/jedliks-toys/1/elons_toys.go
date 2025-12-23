package elon

import "fmt"

// TODO: define the 'Drive()' method
func (car *Car) Drive() {
	var newBattery int = car.battery - car.batteryDrain
	if newBattery >= 0 {
		car.distance += car.speed
		car.battery = newBattery
	}

}

// TODO: define the 'DisplayDistance() string' method
func (car Car) DisplayDistance() string {
	return "Driven " + fmt.Sprint(car.distance) + " meters"
}

// TODO: define the 'DisplayBattery() string' method
func (car Car) DisplayBattery() string {
	return "Battery at " + fmt.Sprint(car.battery) + "%"
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (car Car) CanFinish(trackDistance int) bool {
	var maxDistance int = (car.battery / car.batteryDrain) * car.speed
	return trackDistance <= maxDistance
}
