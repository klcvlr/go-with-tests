package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const sleep = "sleep"
const write = "write"

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

type SpyCountDownOperations struct {
	Calls []string
}

func (s *SpyCountDownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountDownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return len(p), nil
}

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		expected := "3\n2\n1\nGo!"
		spySleeper := &SpyCountDownOperations{}

		Countdown(buffer, spySleeper)

		actualOutput := buffer.String()
		if actualOutput != expected {
			t.Errorf("Expected %q but got %q", expected, actualOutput)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleeperPrinter := &SpyCountDownOperations{}
		expectedSleepCalls := []string{write, sleep, write, sleep, write, sleep, write}

		Countdown(spySleeperPrinter, spySleeperPrinter)

		if !reflect.DeepEqual(spySleeperPrinter.Calls, expectedSleepCalls) {
			t.Errorf("Expected %s 'sleep calls' but got %s", expectedSleepCalls, spySleeperPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := SpyTime{durationSlept: sleepTime}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}

	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
