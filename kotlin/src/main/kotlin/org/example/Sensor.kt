package org.example

import kotlin.random.Random

private const val OFFSET = 16.0

class Sensor {

    fun popNextPressurePsiValue(): Double {
        val pressureTelemetryValue = samplePressure()
        return OFFSET + pressureTelemetryValue
    }

    private fun samplePressure(): Double {
        return 6 * Random.nextDouble() * Random.nextDouble()
    }
}