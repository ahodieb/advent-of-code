package ansi

import "fmt"

func Red(s string) string {
	return fmt.Sprintf("\u001b[31m%s\u001b[0m", s)
}

func RedBG(s string) string {
	return fmt.Sprintf("\u001b[41m%s\u001b[0m", s)
}

func Green(s string) string {
	return fmt.Sprintf("\u001b[32m%s\u001b[0m", s)
}

func GreenBG(s string) string {
	return fmt.Sprintf("\u001b[42m%s\u001b[0m", s)
}

func YellowBG(s string) string {
	return fmt.Sprintf("\u001b[43m%s\u001b[0m", s)
}

func WhiteBG(s string) string {
	return fmt.Sprintf("\u001b[47m%s\u001b[0m", s)
}

func Underline(s string) string {
	return fmt.Sprintf("\u001b[4m%s\u001b[0m", s)
}
