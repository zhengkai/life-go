package life

import (
	"time"
)

// Sleep ...
func Sleep(n int) error {

	if Stop {
		return ErrStop
	}

	select {
	case <-CTX.Done():
	case <-time.After(time.Duration(n) * time.Second):
	}
	if Stop {
		return ErrStop
	}
	return nil
}

// SleepFrom ...
func SleepFrom(t time.Time, n int) (err error) {

	if Stop {
		return ErrStop
	}

	d := t.Add(time.Duration(n) * time.Second).Sub(time.Now())
	// fmt.Println(`SleepFrom`, t, n, d)
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
