package org.example

private const val LOW_PRESSURE_THRESHOLD = 17.0
private const val HIGH_PRESSURE_THRESHOLD = 21.0

class Alarm {

    private val sensor = Sensor()
    private var alarmOn = false
    private var alarmCount = 0

    fun check() {
        val psiPressureValue = sensor.popNextPressurePsiValue()

        if (psiPressureValue < LOW_PRESSURE_THRESHOLD || psiPressureValue > HIGH_PRESSURE_THRESHOLD) {
            alarmOn = true
            alarmCount++
        }
    }

    fun isAlarmOn(): Boolean {
        return alarmOn
    }
}