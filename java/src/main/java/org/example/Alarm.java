package org.example;

public class Alarm {

    private static final double LOW_PRESSURE_THRESHOLD = 17;
    private static final double HIGH_PRESSURE_THRESHOLD = 21;

    Sensor sensor = new Sensor();

    boolean alarmOn = false;
    private long alarmCount = 0;


    public void check()
    {
        double psiPressureValue = sensor.popNextPressurePsiValue();

        if (psiPressureValue < LOW_PRESSURE_THRESHOLD || HIGH_PRESSURE_THRESHOLD < psiPressureValue)
        {
            alarmOn = true;
            alarmCount += 1;
        }
    }

    public boolean isAlarmOn()
    {
       return alarmOn;
    }
}
