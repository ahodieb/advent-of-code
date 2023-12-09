package input

import (
	"bufio"
	"os"
	"strconv"
)

func MustFromPath(path string) *Input {
	in, err := FromPath(path)
	if err != nil {
		panic(err)
	}

	return in
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

func (i *Input) Scan() bool { return i.scanner.Scan() }

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
