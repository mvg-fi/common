package logger

import (
	"fmt"
	"log"
	"regexp"
)

const (
	ERROR   = 1
	INFO    = 2
	VERBOSE = 3
	DEBUG   = 7
)

var (
	level  int
	filter *regexp.Regexp
)

func SetLevel(l int) {
	level = l
}

func SetFilter(pattern string) error {
	if pattern == "" {
		return nil
	}
	// https://github.com/google/re2/wiki/Syntax
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}
	filter = reg
	return nil
}

func Errorf(format string, v ...any) {
	printfAtLevel(ERROR, format, v...)
}

func Println(v ...any) {
	if level >= INFO {
		log.Println(v...)
	}
}

func Printf(format string, v ...any) {
	if level >= INFO {
		log.Printf(format, v...)
	}
}

func Verbosef(format string, v ...any) {
	printfAtLevel(VERBOSE, format, v...)
}

func Debugf(format string, v ...any) {
	printfAtLevel(DEBUG, format, v...)
}

func printfAtLevel(l int, format string, v ...any) {
	if level < l {
		return
	}
	out := filterOutput(format, v...)
	if out == "" {
		return
	}
	log.Print(out)
}

func filterOutput(format string, v ...any) string {
	out := fmt.Sprintf(format, v...)
	if filter == nil || filter.MatchString(out) {
		return out
	}
	return ""
}
