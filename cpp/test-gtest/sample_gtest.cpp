#include <gtest/gtest.h>

#include "alarm.h"

using namespace std;

TEST(Alarm, AlarmIsOffByDefault) {
    auto alarm = Alarm();
    ASSERT_EQ(false, alarm.isAlarmOn());
}