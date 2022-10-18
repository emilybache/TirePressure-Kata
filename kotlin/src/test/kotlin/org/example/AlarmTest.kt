package org.example

import org.junit.jupiter.api.Assertions.assertFalse
import org.junit.jupiter.api.Test

class AlarmTest {

    @Test
    fun alarmIsOffByDefault() {
        val alarm = Alarm()
        assertFalse(alarm.isAlarmOn())
    }

}