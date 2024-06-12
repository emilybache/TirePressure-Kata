package tirepressure

type Alarm struct {
	lowPressureThreshold  float64
	highPressureThreshold float64

	sensor *Sensor

	alarmOn bool
}

func (a Alarm) check() {
	psiPressureValue := a.sensor.popNextPressurePsiValue()

	if psiPressureValue < a.lowPressureThreshold || a.highPressureThreshold < psiPressureValue {
		a.alarmOn = true
	}
}

func (a Alarm) isAlarmOn() bool {
	return a.alarmOn
}

func NewAlarm() *Alarm {
	return &Alarm{lowPressureThreshold: 17, highPressureThreshold: 21, sensor: NewSensor()}
}
