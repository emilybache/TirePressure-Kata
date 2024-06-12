package tirepressure

import "math/rand/v2"

type Sensor struct {
	offset float64
}

func (s Sensor) popNextPressurePsiValue() float64 {
	pressureTelemetryValue := s.samplePressure()
	return s.offset + pressureTelemetryValue
}

func (s Sensor) samplePressure() float64 {
	return 6 * rand.Float64() * rand.Float64()
}

func NewSensor() *Sensor {
	return &Sensor{offset: 16}
}
