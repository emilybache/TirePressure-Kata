#include <gtest/gtest.h>

#include "alarm.h"

using namespace std;

  TEST(Alarm, Foo) {
    auto alarm = new Alarm();
    ASSERT_EQ(false, alarm->isAlarmOn());
  }