
from src.alarm import Alarm


def test_alarm_is_off_by_default():
    alarm = Alarm()
    assert not alarm.is_alarm_on
