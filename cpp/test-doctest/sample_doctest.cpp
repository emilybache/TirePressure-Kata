#include "ApprovalTests.hpp"
#include "doctest/doctest.h"
#include "alarm.h"


TEST_CASE("Alarm is off by default") {
    auto alarm = Alarm();
    REQUIRE(!alarm.isAlarmOn());
}



