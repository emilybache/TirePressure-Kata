using System;
using System.Collections.Generic;
using System.Text;
using Xunit;

namespace TirePressure
{
    public class AlarmTest
    {
        [Fact]
        public void AlarmIsOffByDefault()
        {
            Alarm alarm = new Alarm();
            Assert.False(alarm.AlarmOn);
        }
    }
}