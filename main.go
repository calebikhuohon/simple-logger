package main

import (
	"fmt"
	"io"
	"os"
)

//fmt to show system output
//data passed as bytes

//take in a value with type interface{}
//log that value in STDOUT

//Output to std out or err

type Logger struct {
	out io.Writer
	buf []byte
}
func main()  {
	str := "testing"
	var b []byte
	log := Logger{
		out: os.Stdout,
		buf: b,
	}

	log.log(str)

}

func (l *Logger) log(v interface{})  {
	l.Output(fmt.Sprint(v))
}

func (l *Logger) Output(s string)  error {
	l.buf = append(l.buf, s...)
	if len(s) == 0 || s[len(s) -1] != '\n' {
		l.buf = append(l.buf, '\n')
	}
	_, err := l.out.Write(l.buf)
	return err
}

func (l *Logger) SetOutput(out io.Writer)  {
	//take in a Writer
	l.out = out
}
