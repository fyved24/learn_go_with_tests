package mocking

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}
type ConfigurableSleeper struct {
	Duration time.Duration
}

func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.Duration)
}
func Countdown(w io.Writer, sleeper Sleeper) {

	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()

		_, err := fmt.Fprintln(w, i)
		if err != nil {
			return
		}

	}
	sleeper.Sleep()
	fmt.Fprint(w, finalWord)

}
