package routinue

import "time"

// Routine is
type Routine interface {
	Close()
	Wake()
	Wait() (err error)
}

type goroutine struct {
	wake     chan bool
	exit     chan error
	interval time.Duration
	callback func() error
}

// NewRoutine is
func NewRoutine(interval time.Duration, cb func() error) Routine {
	gr := &goroutine{exit: make(chan error, 1), wake: make(chan bool, 1), interval: interval, callback: cb}

	go gr.loop()

	return gr
}

func (r *goroutine) Close() {
	r.exit <- nil

	close(r.wake)
	close(r.exit)
}

func (r *goroutine) Wake() {
	r.wake <- true
}

func (r *goroutine) Wait() (err error) {
	return <-r.exit
}

func (r *goroutine) loop() {
	for {
		select {
		case <-r.exit:
			return
		case <-r.wake:
		case <-time.After(r.interval):
		}

		if err := r.callback(); nil != err {
			r.exit <- err
			return
		}
	}
}
