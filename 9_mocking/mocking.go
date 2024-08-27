package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

// ? Instead of using bytes.Buffer, use io.Writer for general purpose
func Countdown(out io.Writer, sleeper Sleeper) {
	// ? Since we're not confident about sufficient confidence, we'll break the code
	// // fmt.Fprint(out, "3")
	// for i := countdownStart; i > 0; i-- {
	// 	// ? Need to have newline, use Fprintln
	// 	fmt.Fprintln(out, i)

	// 	// ! This will slow down test, we need to mock this
	// 	// time.Sleep(1 * time.Second)

	// 	// ? Now we will use the sleeper interface
	// 	sleeper.Sleep()
	// }
	// // ? Last line contains no newline, so we use Fprint
	// fmt.Fprint(out, finalWord)

	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
	}

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}

	fmt.Fprint(out, finalWord)
}

// ? We need to DI the sleep
type Sleeper interface {
	Sleep()
}

// ? This will be the "mock"
type SpySleeper struct {
	Calls int
}

// ? This will be the "mock"
func (s *SpySleeper) Sleep() {
	s.Calls++
}

// ?? Since we're breaking down the operation (sleep and write)
// ?? Now we need to make Spy for each operation
const (
	write = "write"
	sleep = "sleep"
)

type SpyCountdownOperations struct {
	Calls []string
}

// ? Mock the sleep operation
func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// ? Mock the write operation
func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

// ? Now we need to create "real" Sleeper
// ?? Since we're using ConfigurableSleeper, we don't need to use DefaultSleeper anymore
// type DefaultSleeper struct{}

// func (d *DefaultSleeper) Sleep() {
// 	time.Sleep(1 * time.Second)
// }

// ? Now we want to make Sleeper to be configurable
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	// ? Now when running countdown, we will use default sleeper
	Countdown(os.Stdout, sleeper)
}
