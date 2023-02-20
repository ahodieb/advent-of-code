package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type SingleIncrementCounter struct {
	counter   int
	lastValue int
}

func NewSingleIncrementCounter(initialValue int) SingleIncrementCounter {
	return SingleIncrementCounter{lastValue: initialValue}
}

func (c *SingleIncrementCounter) Counter() int {
	return c.counter
}

func (c *SingleIncrementCounter) Count(n int) {
	if n > c.lastValue {
		c.counter += 1
	}

	c.lastValue = n
}

type WindowIncrementCounter struct {
	window  []int
	counter int
	wSize   int
}

func NewWindowIncrementCounter(windowSize int) WindowIncrementCounter {
	return WindowIncrementCounter{
		wSize: windowSize,
	}
}

func (c *WindowIncrementCounter) Counter() int {
	return c.counter
}

func (c *WindowIncrementCounter) Count(n int) {
	c.window = append(c.window, n)

	if len(c.window) <= c.wSize {
		return
	}

	p := c.window[len(c.window)-1-c.wSize]
	sum := 0
	for _, i := range c.window[len(c.window)-1-c.wSize : len(c.window)-1] {
		sum += i
	}

	if sum+n > sum+p {
		c.counter += 1
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Scan()
	first := atoi(s.Text())

	c1 := NewSingleIncrementCounter(first)

	c2 := NewWindowIncrementCounter(3)
	c2.Count(first)

	for s.Scan() {
		c1.Count(atoi(s.Text()))
		c2.Count(atoi(s.Text()))
	}

	fmt.Println(c1.Counter())
	fmt.Println(c2.Counter())
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
