package server

import (
	"context"
	"fmt"
	"log"
	"os"
)

var debug *log.Logger
var info *log.Logger
var warn *log.Logger
var fatal *log.Logger

func level(level, tag string) string {
	return fmt.Sprintf("%s [%s] ", level, tag)
}

func initLogger(tag string) {
	f := log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC | log.Lshortfile

	debug = log.New(os.Stdout, level("DEBG", tag), f)
	info = log.New(os.Stdout, level("INFO", tag), f)
	warn = log.New(os.Stdout, level("WARN", tag), f)
	fatal = log.New(os.Stdout, level("FATAL", tag), f)
}

func Debug(c context.Context, m string) {
	debug.Printf("%s%s", contextify(c), m)
}

func Debugf(c context.Context, m string, v ...interface{}) {
	Debug(c, fmt.Sprintf(m, v...))
}

func Info(c context.Context, m string) {
	info.Printf("%s%s", contextify(c), m)
}

func Infof(c context.Context, m string, v ...interface{}) {
	Info(c, fmt.Sprintf(m, v...))
}

func Warn(c context.Context, m string) {
	warn.Printf("%s%s", contextify(c), m)
}

func Warnf(c context.Context, m string, v ...interface{}) {
	Warn(c, fmt.Sprintf(m, v...))
}

func Fatal(c context.Context, m string) {
	fatal.Fatalf("%s%s", contextify(c), m)
}

func Fatalf(c context.Context, m string, v ...interface{}) {
	Fatal(c, fmt.Sprintf(m, v...))
}

func contextify(c context.Context) string {
	id := c.Value("correlationId")
	prefix := ""
	if id != nil {
		prefix = fmt.Sprintf("%s ", id)
	}

	path := c.Value("path")

	if path != nil {
		prefix = fmt.Sprintf("%s%s ", prefix, path)
	}

	return prefix
}
