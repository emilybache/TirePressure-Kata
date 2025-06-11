package tirepressure_test

import (
	"testing"

	tirepressure "github.com/emilybache/TirePressure-Kata/go"
	"github.com/stretchr/testify/assert"
)

func TestAlarm(t *testing.T) {
	t.Run("alarm is off by default", func(t *testing.T) {
		alarm := tirepressure.NewAlarm()

		assert.False(t, alarm.IsAlarmOn())
	})
}
