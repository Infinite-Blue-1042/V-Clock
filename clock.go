package clock
import (
	"sort"
	"sync"
	"time"
)
type Clock interface {
	After(d time.Duration) <-chan time.Time
	AfterFunc(d time.Duration, f func()) *Timer
	Now() time.Time
	Since(t time.Time) time.Duration
	Sleep(d time.Duration)
	Tick(d time.Duration) <-chan time.Time
	Ticker(d time.Duration) *Ticker
	Timer(d time.Duration) *Timer
}
func New() Clock {
	return &clock{}
}
type clock struct{}
func (c *clock) After(d time.Duration) <-chan time.Time { return time.After(d) }
func (c *clock) AfterFunc(d time.Duration, f func()) *Timer {
	return &Timer{timer: time.AfterFunc(d, f)}
}
func (c *clock) Now() time.Time { return time.Now() }
func (c *clock) Since(t time.Time) time.Duration { return time.Since(t) }
func (c *clock) Sleep(d time.Duration) { time.Sleep(d) }
func (c *clock) Tick(d time.Duration) <-chan time.Time { return time.Tick(d) }
func (c *clock) Ticker(d time.Duration) *Ticker {
	t := time.NewTicker(d)
	return &Ticker{C: t.C, ticker: t}
}
func (c *clock) Timer(d time.Duration) *Timer {
	t := time.NewTimer(d)
	return &Timer{C: t.C, timer: t}
}
