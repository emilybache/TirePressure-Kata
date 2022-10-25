

#include "sensor.h"
#include "iostream"

#include <cstdlib>

Sensor::Sensor()
{
    srand(time(NULL));
}

double Sensor::popNextPressurePsiValue()
{
    // placeholder implementation that simulates a real sensor in a real tire
    float offset = (float) (rand() % 8);
    //std::cout << "offset: " << offset << "\n";
    double pressure = 15 + offset;
    return pressure;
}