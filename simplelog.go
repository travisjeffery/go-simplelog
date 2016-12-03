package simplelog

import (
	"fmt"
	"io"
	"sync"
	"time"
)

type Level int

const (
	DEBUG Level = iota
	INFO
)

var levels = map[Level]string{
	DEBUG: "DEBUG",
	INFO:  "INFO",
}

type Logger struct {
	Writer io.Writer
	Level  Level
	Prefix string
	sync.Mutex
}

func New(writer io.Writer, level Level, prefix string) *Logger {
	return &Logger{
		Writer: writer,
		Level:  level,
		Prefix: prefix,
	}
}

func (l *Logger) log(level Level, msg string, args ...interface{}) error {
	l.Lock()
	defer l.Unlock()
	if l.Level > level {
		return nil
	}
	ts := time.Now().Format("2006-01-02 15:04:05")
	f := fmt.Sprintf("%s [%s] %s: %s\n", ts, levels[level], l.Prefix, msg)
	_, err := fmt.Fprintf(l.Writer, f, args...)
	return err
}

func (l *Logger) Info(msg string, args ...interface{}) error {
	return l.log(INFO, msg, args...)
}

func (l *Logger) Debug(msg string, args ...interface{}) error {
	return l.log(DEBUG, msg, args...)
}
