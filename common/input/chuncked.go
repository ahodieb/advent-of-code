package input

type Chunked struct {
	input     *Input
	chunkSize int
	chunk     []string
}

func (c *Chunked) Close() { c.input.Close() }

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

func ChunkedFromPath(path string, chunk int) (*Chunked, error) {
	in, err := FromPath(path)
	if err != nil {
		return nil, err
	}

	return &Chunked{input: in, chunkSize: chunk}, nil
}
