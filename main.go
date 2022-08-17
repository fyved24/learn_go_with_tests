package main

import (
	"learn_go_with_tests/7.mocking"
	"os"
	"time"
)

func main() {
	sleeper := &mocking.ConfigurableSleeper{Duration: 1 * time.Second}
	mocking.Countdown(os.Stdout, sleeper)
}
