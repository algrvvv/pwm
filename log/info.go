package log

import "fmt"

func Sinfo(a any) string {
	return fmt.Sprintf("\n   %s  %v", info(" INFO "), a)
}

func Info(a any) {
	fmt.Println(Sinfo(a))
}

func Infof(format string, args ...any) {
	Info(fmt.Sprintf(format, args...))
}
