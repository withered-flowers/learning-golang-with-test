package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {

	t.Run("Print 3 to go !", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}

		// ? While we're doing the testing, we will use the SpySleeper instead of DefaultSleeper
		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if spySleeper.Calls != countdownStart {
			t.Errorf("Not enough calls to sleeper, want %d got %d", countdownStart, spySleeper.Calls)
		}
	})

	t.Run("Sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		// ? Now we will use the spySleepPrinter for Sleep and Write operation
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}

	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should sleep for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
