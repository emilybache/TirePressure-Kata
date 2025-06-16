package tirepressure

const (
	lowPressureThreshold  float64 = 17
	highPressureThreshold float64 = 21
)

type Alarm struct {
	sensor  *Sensor
	alarmOn bool
}

func (a *Alarm) Check() {
	psiPressureValue := a.sensor.PopNextPressurePsiValue()

	if psiPressureValue < lowPressureThreshold || highPressureThreshold < psiPressureValue {
		a.alarmOn = true
	}
}

func (a Alarm) IsAlarmOn() bool {
	return a.alarmOn
}

func NewAlarm() *Alarm {
	return &Alarm{sensor: NewSensor()}
}
