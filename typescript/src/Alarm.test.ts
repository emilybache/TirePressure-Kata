import Alarm from './Alarm';

describe('Tire Pressure Alarm', () => {
  it('should be off by default', () => {
    const alarm = new Alarm();

    expect(alarm.isOn()).toBeFalsy();
  });
});
