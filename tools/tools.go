package tools

import "fmt"

type Logger interface {
	Log(string)
}

type ConsoleLogger struct {
	Format  string
	Spliter string
}

func (cl ConsoleLogger) Log(s string) {
	fmt.Printf(cl.Format, s, cl.Spliter)
}
