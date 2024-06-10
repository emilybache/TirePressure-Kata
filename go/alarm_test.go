package tirepressure

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlarm(t *testing.T) {
	t.Run("alarm is off by default", func(t *testing.T) {
		alarm := NewAlarm()

		assert.False(t, alarm.isAlarmOn())
	})
}
