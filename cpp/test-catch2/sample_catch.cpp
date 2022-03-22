#include "ApprovalTests.hpp"
#include "catch2/catch.hpp"
#include "alarm.h"


TEST_CASE("Alarm", "[TirePressureMonitoring]") {
    auto alarm = new Alarm();
    REQUIRE(!alarm->isAlarmOn());
}



