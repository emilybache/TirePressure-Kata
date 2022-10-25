#include "ApprovalTests.hpp"
#include "catch2/catch.hpp"
#include "alarm.h"


TEST_CASE("Alarm", "is off by default") {
    auto alarm = Alarm();
    REQUIRE(!alarm.isAlarmOn());
}



