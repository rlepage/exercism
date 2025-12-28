package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	Hours   int
	Minutes int
}

func (c *Clock) NormalizeClock() {
	for c.Minutes < 0 {
		c.Minutes += 60
		c.Hours--
	}
	c.Hours += int(c.Minutes / 60)
	c.Minutes %= 60

	for c.Hours < 0 {
		c.Hours += 24
	}
	c.Hours %= 24
}

func New(h, m int) Clock {
	c := Clock{Hours: h, Minutes: m}
	c.NormalizeClock()
	return c
}

func (c Clock) Add(m int) Clock {
	c.Minutes += m
	c.NormalizeClock()
	return c
}

func (c Clock) Subtract(m int) Clock {
	c.Minutes -= m
	c.NormalizeClock()
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hours, c.Minutes)
}
