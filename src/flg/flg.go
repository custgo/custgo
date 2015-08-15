package flg

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// File LoGger

var defaultTimeFormat = "2006-01-02 15:04:05"

type Logger struct {
	mu     sync.Mutex // ensures atomic writes; protects the following fields
	prefix string     // prefix to write at beginning of each line
	ftime  string     // properties
	out    io.Writer  // destination for output
	buf    []byte     // for accumulating text to write
}

func New(filename string) *Logger {
	if "" == filename {
		return &Logger{ftime: defaultTimeFormat, out: os.Stderr}
	}
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0644)
	if nil != err {
		fmt.Println(err)
		return &Logger{ftime: defaultTimeFormat, out: os.Stderr}
	}
	return &Logger{ftime: defaultTimeFormat, out: file}
}

func (l *Logger) SetFilename(filename string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if f, ok := l.out.(*os.File); ok {
		f.Close()
	}
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0644)
	if nil != err {
		fmt.Println(err)
		l.out = os.Stderr
		return
	}
	l.out = file
}

func (l *Logger) SetWriter(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.out = w
}

func (l *Logger) Prefix() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.prefix
}

func (l *Logger) SetPrefix(prefix string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.prefix = prefix
}

func (l *Logger) TimeFormat() string {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.ftime
}

func (l *Logger) SetTimeFormat(format string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.ftime = format
}

func (l *Logger) Write(msg string) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.buf = l.buf[:0]
	if "" != l.ftime {
		l.buf = append(l.buf, time.Now().Format(l.ftime)...)
		l.buf = append(l.buf, ' ')
	}
	if "" != l.prefix {
		l.buf = append(l.buf, l.prefix...)
	}
	l.buf = append(l.buf, msg...)
	_, err := l.out.Write(l.buf)
	return err
}

func (l *Logger) Print(v ...interface{}) {
	l.Write(fmt.Sprint(v...))
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.Write(fmt.Sprintf(format, v...))
}

func (l *Logger) Println(v ...interface{}) {
	l.Write(fmt.Sprintln(v...))
}

func (l *Logger) Fatal(v ...interface{}) {
	l.Write(fmt.Sprint(v...))
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Write(fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (l *Logger) Fatalln(v ...interface{}) {
	l.Write(fmt.Sprintln(v...))
	os.Exit(1)
}

func (l *Logger) Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	l.Write(s)
	panic(s)
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.Write(s)
	panic(s)
}

func (l *Logger) Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	l.Write(s)
	panic(s)
}

var std = New("")

func SetWriter(w io.Writer) {
	std.SetWriter(w)
}

func Write(msg string) error {
	return std.Write(msg)
}

func TimeFormat() string {
	return std.TimeFormat()
}

func SetTimeFormat(format string) {
	std.SetTimeFormat(format)
}

func Prefix() string {
	return std.Prefix()
}

func SetPrefix(prefix string) {
	std.SetPrefix(prefix)
}

func SetFilename(filename string) {
	std.SetFilename(filename)
}

func Print(v ...interface{}) {
	std.Print(v...)
}

func Printf(format string, v ...interface{}) {
	std.Printf(format, v...)
}

func Println(v ...interface{}) {
	std.Println(v...)
}

func Fatal(v ...interface{}) {
	std.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	std.Fatalf(format, v...)
}

func Fatalln(v ...interface{}) {
	std.Fatalln(v...)
}

func Panic(v ...interface{}) {
	std.Panic(v...)
}

func Panicf(format string, v ...interface{}) {
	std.Panicf(format, v...)
}

func Panicln(v ...interface{}) {
	std.Panicln(v...)
}
