

#ifndef TIREPRESSURE_ALARM_H
#define TIREPRESSURE_ALARM_H

#include "sensor.h"

class Alarm
{
public:
    Alarm();
    void check(const std::function<double(Alarm *)> &getPressureLambda);
    bool isAlarmOn();

protected:
    Sensor m_sensor;
    double m_lowPressureThreshold;
    double m_highPressureThreshold;
    bool m_alarmOn;

    double getPressure();
};


#endif //TIREPRESSURE_ALARM_H
