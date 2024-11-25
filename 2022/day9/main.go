package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ahodieb/brute/input"
)

func main() {
	in := input.FromPath("day9/input.txt")
	defer in.Close()

	rope := Rope{Knots: make([]Knot, 10)}
	visited := make(map[string]struct{})
	visitedCount := 0

	for in.Scan() {
		move := strings.Split(in.Text(), " ")
		direction := move[0]
		steps, _ := strconv.Atoi(move[1])

		for i := 0; i < steps; i++ {
			rope.MoveHead(direction)

			if _, ok := visited[rope.Tail().String()]; !ok {
				visited[rope.Tail().String()] = struct{}{}
				visitedCount += 1
			}
		}
	}

	fmt.Println("Tail moved", visitedCount)
}

type Rope struct {
	Knots []Knot
}

func (r *Rope) MoveHead(direction string) {
	r.Knots[0].Move(direction)

	for i := 1; i < len(r.Knots); i++ {
		r.Knots[i].MoveTowards(r.Knots[i-1])
	}
}

func (r *Rope) Tail() Knot {
	return r.Knots[len(r.Knots)-1]
}

type Knot struct {
	X int
	Y int
}

func (k *Knot) Move(direction string) {
	switch direction {
	case "R":
		k.X += 1
	case "L":
		k.X -= 1
	case "U":
		k.Y += 1
	case "D":
		k.Y -= 1
	}
}

func (k Knot) String() string {
	return fmt.Sprintf("x=%dy=%d", k.X, k.Y)
}

func (k *Knot) IsTouching(kk Knot) bool {
	// ....
	// .HT.  -> Y == Y , X == X+1
	// ....

	// ....
	// TH..  -> Y == Y , X == X-1
	// ....

	return (k.Y == kk.Y && (k.X == kk.X || k.X+1 == kk.X || k.X-1 == kk.X)) ||

		// .T..
		// .H..  -> X == X , Y == Y+1
		// ....

		// ....
		// .H..  -> X == X , Y == Y01
		// .T..

		(k.X == kk.X && (k.Y+1 == kk.Y || k.Y-1 == kk.Y)) ||

		// ..T.
		// .H..  -> X == X+1, Y == Y+1
		// ....

		// ....
		// .H..  -> X == X+1, Y == Y-1
		// ..T.

		// ....
		// .H..  -> X == X-1, Y == Y-1
		// T...

		// T...
		// .H..  -> X == X-1, Y == Y+1
		// ....

		(k.X+1 == kk.X && k.Y+1 == kk.Y) ||
		(k.X+1 == kk.X && k.Y-1 == kk.Y) ||
		(k.X-1 == kk.X && k.Y-1 == kk.Y) ||
		(k.X-1 == kk.X && k.Y+1 == kk.Y)
}

func (k *Knot) MoveTowards(kk Knot) {
	if k.IsTouching(kk) {
		return
	}

	if kk.X > k.X {
		k.X += 1
	} else if kk.X < k.X {
		k.X -= 1
	}

	if kk.Y > k.Y {
		k.Y += 1
	} else if kk.Y < k.Y {
		k.Y -= 1
	}
}
