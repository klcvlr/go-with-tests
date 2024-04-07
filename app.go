package main

import (
	"go-with-test/mocking"
	"os"
	"time"
)

func main() {
	sleeper := &mocking.ConfigurableSleeper{Duration: 500 * time.Millisecond, SleepFunc: time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
}
