package input

import (
	"bufio"
	"os"
	"strconv"
)

func ChunkedFromPath(path string, chunk int) (*Chunked, error) {
	in, err := FromPath(path)
	if err != nil {
		return nil, err
	}

	return &Chunked{input: in, chunkSize: chunk}, nil
}

func FromPath(path string) (*Input, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return &Input{
		scanner: bufio.NewScanner(f),
		closer:  func() { _ = f.Close() },
	}, nil
}

type Input struct {
	scanner *bufio.Scanner
	closer  func()
}

func (i *Input) Close() { i.closer() }

func (i *Input) Scan() bool {
	return i.scanner.Scan()
}

func (i *Input) ScanAll() []string {
	var lines []string
	for i.Scan() {
		lines = append(lines, i.Text())
	}

	return lines
}

func (i *Input) Text() string {
	return i.scanner.Text()
}

func (i *Input) Int() int {
	n, _ := strconv.Atoi(i.scanner.Text())
	return n
}

type Chunked struct {
	input     *Input
	chunkSize int
	chunk     []string
}

func (c *Chunked) Close()          { c.input.Close() }
func (c *Chunked) Chunk() []string { return c.chunk }

func (c *Chunked) Scan() bool {
	c.chunk = nil

	if c.chunkSize < 2 {
		next := c.input.Scan()
		if next {
			c.chunk = []string{c.input.Text()}
		}
		return next
	}

	for i := 0; i < c.chunkSize; i++ {
		next := c.input.Scan()
		if !next {
			return len(c.chunk) > 0
		}

		c.chunk = append(c.chunk, c.input.Text())
	}

	return true
}
