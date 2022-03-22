#include "ApprovalTests.hpp"
#include "doctest/doctest.h"
#include "alarm.h"


TEST_CASE("Alarm [TirePressureMonitoring]") {
    auto alarm = new Alarm();
    REQUIRE(!alarm->isAlarmOn());
}



