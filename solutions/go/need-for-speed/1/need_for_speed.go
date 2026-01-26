package speed

// TODO: define the 'Car' type struct
type Car struct {
    battery int
    batteryDrain int
    speed int
    distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	car := Car {
        battery: 100,
        speed: speed,
        batteryDrain: batteryDrain,
    }
    return car
}

// TODO: define the 'Track' type struct
type Track struct {
    distance int
}
// NewTrack creates a new track
func NewTrack(distance int) Track {
	track := Track {
        distance: distance,
    }
    return track
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
    // Check if we have enough battery to drive.
    usedBattery := car.battery - car.batteryDrain
    if usedBattery < 0 {
        return car
    }
    distanceMoved := car.distance + car.speed
	droveCar := Car{
        battery: usedBattery,
		batteryDrain: car.batteryDrain,
		speed: car.speed,
		distance: distanceMoved,
    }
    return droveCar
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
    carCanTravel := 0
    // Drive the car until batteries discharged.
    for car.battery > car.batteryDrain {
        car = Drive(car)
        carCanTravel = car.distance
    }
    // Compare the distance travelled and track distance.
    if carCanTravel >= track.distance {
        return true
    }
    return false
}
