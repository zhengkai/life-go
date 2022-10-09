package life

import (
	"time"
)

// Sleep 等待 n 秒，如果有信号则提前退出
func Sleep(n int) error {
	d := time.Duration(n) * time.Second
	return SleepDur(d)
}

// SleepDur 等待指定时间，如果有信号则提前退出
func SleepDur(d time.Duration) error {

	if Stop {
		return ErrStop
	}

	select {
	case <-CTX.Done():
	case <-time.After(d):
	}
	if Stop {
		return ErrStop
	}
	return nil
}

// SleepFrom 以某个时间点为准，等待 n 秒，如果有信号则提前退出
func SleepFrom(t time.Time, n int) (err error) {

	d := time.Duration(n) * time.Second
	return SleepDurFrom(t, d)

}

// SleepDurFrom 以某个时间点为准，等待 n 秒，如果有信号则提前退出
func SleepDurFrom(t time.Time, d time.Duration) (err error) {

	if Stop {
		return ErrStop
	}

	d = t.Add(d).Sub(time.Now())
	if d < 0 {
		return
	}

	select {
	case <-CTX.Done():
		err = ErrStop
	case <-time.After(d):
	}
	if Stop {
		err = ErrStop
	}
	return
}
