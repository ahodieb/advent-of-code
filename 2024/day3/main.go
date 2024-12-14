package main

import (
	"fmt"
	"github.com/ahodieb/brute/input"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var mul = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

var re = regexp.MustCompile(`(mul\((\d{1,3}),(\d{1,3})\))|(do\(\))|(don't\(\))`)
var dont = regexp.MustCompile(`don't\(\).*(do\(\))|$`)

func main() {
	day := "day3"
	//solve(fmt.Sprintf("2024/%s/input-small.txt", day))    // 161
	//solve2(fmt.Sprintf("2024/%s/input-small-p2.txt", day)) // 48
	solve3(fmt.Sprintf("2024/%s/input.txt", day)) // 174561379 106921067
}

func solve3(p string) {
	in := input.FromPath(p)
	defer in.Close()

	for in.Scan() {
		txt := in.Text()
		mem := txt
		for _, r := range dont.FindAllStringIndex(txt, -1) {
			mem = strings.ReplaceAll(mem, txt[r[0]:r[1]], "")
		}
		for _, m := range mul.FindAllString(mem, -1) {
			fmt.Println(m)
		}
	}
}

func solve2(p string) {
	in := input.FromPath(p)
	defer in.Close()
	sum := 0
	for in.Scan() {
		txt := in.Text()
		//enabled := true
		for i := 0; i < len(txt); {
			fmt.Println(txt[i:])

			index := findFirstIndex(txt[i:], "mul(", "do()", "don't()")
			if index == -1 {
				break
			}

			if strings.HasPrefix(txt[i+index:], "mul(") {
				m := mul.FindStringSubmatch(txt[i+index:])
				if m == nil {
					i += index + 4
					continue
				}
				fmt.Println(m, ">>> ", txt[i+index:])
				a, _ := strconv.Atoi(m[1])
				b, _ := strconv.Atoi(m[2])
				sum += a * b
				i += index + 4
			} else if strings.HasPrefix(txt[i+index:], "do()") {
				i += index + 4
			} else if strings.HasPrefix(txt[i+index:], "don't()") {
				skip := strings.Index(txt[i+index:], "do()")
				if skip == -1 {
					break
				}
				i += index + skip
			} else {
				//fmt.Println("WTF >>> ", txt[i+index:])
				i += 1
			}
		}

	}

	fmt.Println(sum)
}

func solve(p string) {
	in := input.FromPath(p)
	defer in.Close()
	sum := 0
	sumCond := 0
	for in.Scan() {
		//fmt.Println(in.Text())
		tokens := extractRegex(in.Text())
		//tokens := parse(in.Text())
		//fmt.Printf("%+v\n", tokens)
		do := true
		for _, t := range tokens {
			if _, ok := t.(*Do); ok {
				do = true
			} else if _, ok := t.(*Dont); ok {
				do = false
			} else if v, ok := t.(*Mul); ok {
				mm := v.B * v.A
				sum = sum + mm
				if do {
					sumCond = sumCond + mm
				}
			} else {
				fmt.Println("WTF", t)
			}
		}
	}

	fmt.Println(sum, sumCond)
}

func extractRegex(p string) []interface{} {
	matches := re.FindAllStringSubmatch(p, -1)
	var tokens []interface{}
	for _, m := range matches {
		if strings.HasPrefix(m[0], "mul") {
			a, _ := strconv.Atoi(m[2])
			b, _ := strconv.Atoi(m[3])
			tokens = append(tokens, &Mul{a, b})
		} else if strings.HasPrefix(m[0], "don't") {
			tokens = append(tokens, &Dont{})
		} else if strings.HasPrefix(m[0], "do") {
			tokens = append(tokens, &Do{})
		} else {
			fmt.Println(m)
		}
	}
	return tokens
}

func parse(p string) []interface{} {
	var tokens []interface{}
	for i := 0; i < len(p); {
		//fmt.Println(p[i:])
		if strings.HasPrefix(p[i:], "mul(") {
			m := readMul(p[i:])
			if m != nil {
				tokens = append(tokens, m)
				i = i + len(m.String())
			} else {
				i = i + len("mul(")
			}
		} else if strings.HasPrefix(p[i:], "don't()") {
			tokens = append(tokens, &Dont{})
			i = i + len("don't()")
		} else if strings.HasPrefix(p[i:], "do()") {
			tokens = append(tokens, &Do{})
			i = i + len("do()")
		} else {
			nxt := findFirstIndex(p[i:], "mul(", "do(", "don't(")
			if nxt == -1 {
				break
			}
			i = i + nxt
		}
	}

	return tokens
}

type Mul struct {
	A int
	B int
}

func (m *Mul) String() string {
	return fmt.Sprintf("mul(%d,%d)", m.A, m.B)
}

type Do struct{}

func (m *Do) String() string {
	return "do()"
}

type Dont struct{}

func (m *Dont) String() string {
	return "don't()"
}

func readMul(p string) *Mul {
	match := mul.FindStringSubmatch(p)
	if match == nil {
		return nil
	}
	a, _ := strconv.Atoi(match[1])
	b, _ := strconv.Atoi(match[2])
	return &Mul{a, b}
}

func findFirstIndex(s string, tokens ...string) int {
	if len(tokens) == 0 {
		return -1
	}

	var index []int
	for _, t := range tokens {
		i := strings.Index(s, t)
		if i == -1 {
			continue
		}
		index = append(index, i)
	}

	if len(index) == 0 {
		return -1
	}

	slices.Sort(index)
	return index[0]
}
