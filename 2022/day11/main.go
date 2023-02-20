package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/ahodieb/advent-of-code/common/input"
)

func main() {
	in := input.MustFromPath("day11/input.txt")
	defer in.Close()

	var monkeys []Monkey

	numbers := map[int]struct{}{
		1:     {},
		20:    {},
		1000:  {},
		2000:  {},
		3000:  {},
		4000:  {},
		5000:  {},
		6000:  {},
		7000:  {},
		9000:  {},
		10000: {},
	}

	mod := 1
	for in.Scan() {
		monkey := Monkey{Index: len(monkeys)}

		in.Scan()
		items := strings.Split(TrimPrefix(in.Text(), "Starting items:"), ",")
		for _, item := range items {
			v, _ := strconv.Atoi(strings.TrimSpace(item))
			monkey.Items = append(monkey.Items, v)
		}

		in.Scan()
		operation := TrimPrefix(in.Text(), "Operation: new =")
		monkey.OperationStr = operation
		operator := "+"
		if strings.Contains(operation, "*") {
			operator = "*"
		}
		value := strings.TrimSpace(strings.Split(operation, operator)[1])
		if value == "old" {
			if operator == "+" {
				monkey.Operation = Double()
			} else {
				monkey.Operation = Square()
			}
		} else {
			v, _ := strconv.Atoi(value)
			if operator == "+" {
				monkey.Operation = AddValue(v)
			} else {
				monkey.Operation = MulValue(v)
			}
		}

		in.Scan()
		monkey.TestDivisibleBy = Atoi(TrimPrefix(in.Text(), "Test: divisible by"))
		mod *= monkey.TestDivisibleBy

		in.Scan()
		monkey.TestTrue = Atoi(TrimPrefix(in.Text(), "If true: throw to monkey"))

		in.Scan()
		monkey.TestFalse = Atoi(TrimPrefix(in.Text(), "If false: throw to monkey"))

		monkeys = append(monkeys, monkey)
		in.Scan()
	}

	for steps := 0; steps < 10000; steps++ {
		for m := range monkeys {
			//fmt.Printf("Monkey %d:\n", m)

			count := len(monkeys[m].Items)
			monkeys[m].Inspected += count
			for i := 0; i < count; i++ {
				item := monkeys[m].Items[0]
				monkeys[m].Items = monkeys[m].Items[1:]
				//fmt.Printf("  Monkey inspects an item with a worry level of %d.\n", item)

				item = monkeys[m].Operation(item)
				//fmt.Printf("    Worry level changed to %d.\n", item)

				//item = item / 3

				// Used a hint from
				// https://www.reddit.com/r/adventofcode/comments/zih7gf/comment/izrck61/?utm_source=share&utm_medium=web2x&context=3
				item %= mod //
				//fmt.Printf("    Monkey gets bored with item. Worry level is divided by 3 to %d.\n", item)

				nextMonkey := monkeys[m].TestFalse
				if item%monkeys[m].TestDivisibleBy == 0 {
					nextMonkey = monkeys[m].TestTrue
					//fmt.Printf("    Current worry level is divisible by %d.\n", monkeys[m].TestDivisibleBy)
				} else {
					//fmt.Printf("    Current worry level is divisible by %d.\n", monkeys[m].TestDivisibleBy)
				}

				//fmt.Printf("    Item with worry level %d is thrown to monkey %d.\n", item, nextMonkey)
				monkeys[nextMonkey].Items = append(monkeys[nextMonkey].Items, item)
			}
		}

		if _, ok := numbers[steps+1]; ok {
			fmt.Println("Round:", steps+1)
			for i := range monkeys {
				fmt.Printf("Monkey %d inspected items %d times.\n", i, monkeys[i].Inspected)
			}
			//_, _ = fmt.Scanln()
		}

	}

	var inspectedTimes []int
	for i := range monkeys {
		fmt.Println(monkeys[i].Result())
		inspectedTimes = append(inspectedTimes, monkeys[i].Inspected)
	}

	sort.Ints(inspectedTimes)
	fmt.Println(inspectedTimes[len(inspectedTimes)-1] * inspectedTimes[len(inspectedTimes)-2])
}

type Monkey struct {
	Items           []int
	Operation       func(old int) int
	TestDivisibleBy int
	OperationStr    string
	Index           int
	TestTrue        int
	TestFalse       int

	Inspected int
}

func (m *Monkey) String() string {
	return fmt.Sprintf("Monkey %d op=%-9s test=%d t=%d f=%d %v", m.Index, m.OperationStr, m.TestDivisibleBy, m.TestTrue, m.TestFalse, m.Items)
}

func (m *Monkey) Result() string {
	return fmt.Sprintf("Monkey %d [%02d] %v", m.Index, m.Inspected,
		m.Items)
}

func AddValue(x int) func(int) int {
	return func(old int) int {
		return old + x
	}
}

func MulValue(x int) func(int) int {
	return func(old int) int {
		return old * x
	}
}

func Double() func(int) int {
	return func(old int) int {
		return old + old
	}
}

func Square() func(int) int {
	return func(old int) int {
		return old * old
	}
}

func TrimPrefix(s, prefix string) string {
	return strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(s), prefix))
}

func Atoi(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return v
}
