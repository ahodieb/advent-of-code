package grid

type Iterator[T any] struct {
	g              *Grid[T]
	firstIteration bool
	curPos         Position
}

type Cell[T any] struct {
	Value    T
	Position Position
}

func (i *Iterator[T]) Next() bool {
	size := i.g.Size()

	if i.firstIteration == false {
		i.firstIteration = true
		return size.Rows != 0 && size.Columns != 0
	}

	i.curPos.Column += 1
	if i.curPos.Column >= size.Columns {
		i.curPos.Column = 0
		i.curPos.Row += 1
	}

	return i.curPos.Row < size.Rows
}

func (i *Iterator[T]) Position() Position {
	return i.curPos
}

func (i *Iterator[T]) Value() T {
	return i.g.Value(i.curPos)
}

func (i *Iterator[T]) Cell() Cell[T] {
	return i.g.Cell(i.curPos)
}
