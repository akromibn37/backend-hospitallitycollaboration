package logging

import "time"

type Timelog struct {
	TimeIn    time.Time
	TimeOut   time.Time
	TimeInDb  time.Time
	TimeOutDb time.Time
	DiffE2e   time.Duration
	DiffDB    time.Duration
	FuncNm    string
}

func (t *Timelog) GetDiff() {
	t.DiffE2e = t.TimeOut.Sub(t.TimeIn)
	t.DiffDB = t.TimeOutDb.Sub(t.TimeInDb)
}
