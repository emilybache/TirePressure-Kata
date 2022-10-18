package org.example;

import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertFalse;

public class AlarmTest {

    @Test
    public void alarmIsOffByDefault()
    {
        Alarm alarm = new Alarm();
        Assertions.assertFalse(alarm.isAlarmOn());
    }

}
