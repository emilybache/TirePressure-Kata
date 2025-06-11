package tirepressure

import "math/rand/v2"

const offset float64 = 16

type Sensor struct{}

func (s Sensor) PopNextPressurePsiValue() float64 {
	pressureTelemetryValue := samplePressure()
	return offset + pressureTelemetryValue
}

func samplePressure() float64 {
	return 6 * rand.Float64() * rand.Float64()
}

func NewSensor() *Sensor {
	return &Sensor{}
}
