package mocking

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

const write = "write"
const sleep = "sleep"

type Sleeper interface {
	Sleep()
}
type ConfigurableSleeper struct {
	Duration time.Duration
}

func (o *ConfigurableSleeper) Sleep() {
	time.Sleep(o.Duration)
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
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
