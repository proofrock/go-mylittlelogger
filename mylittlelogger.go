/*
  Copyright (c) 2022-, Germano Rizzo <oss@germanorizzo.it>

  Permission to use, copy, modify, and/or distribute this software for any
  purpose with or without fee is hereby granted, provided that the above
  copyright notice and this permission notice appear in all copies.

  THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
  WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
  MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
  ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
  WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
  ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
  OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
*/

// Version 0.3.0

package mylittlelogger

import (
	"fmt"
	"os"
	"time"
)

const DEBUG = 4
const INFO = 3
const WARN = 2
const ERROR = 1
const FATAL = 0
const NONE = -1
const NOT_EVEN_STDOUT = -2
const NOT_EVEN_STDERR = -3

var ForFatal = func() { os.Exit(1) }
var Level = INFO
var Prefixes = []string{"DEBUG", "INFO", "WARN", "ERR", "FATAL"}
var DateTimeFormat = "2006/01/02 15:04:05"

func line(lvl int, a ...interface{}) string {
	return fmt.Sprintf("%s | %s | %s\n", Prefixes[Level], time.Now().Format(DateTimeFormat), fmt.Sprint(a...))
}

func linef(lvl int, format string, elements ...interface{}) string {
	return line(lvl, fmt.Sprintf(format, elements...))
}

func StdOut(a ...interface{}) {
	if Level > NOT_EVEN_STDOUT {
		fmt.Fprint(os.Stdout, fmt.Sprint(a...), "\n")
	}
}

func StdOutf(format string, elements ...interface{}) {
	if Level > NOT_EVEN_STDOUT {
		fmt.Fprint(os.Stdout, fmt.Sprintf(format, elements...), "\n")
	}
}

func StdOutl(lambda func() string) {
	if Level > NOT_EVEN_STDOUT {
		fmt.Fprint(os.Stdout, lambda(), "\n")
	}
}

func StdErr(a ...interface{}) {
	if Level > NOT_EVEN_STDERR {
		fmt.Fprint(os.Stderr, fmt.Sprint(a...), "\n")
	}
}

func StdErrf(format string, elements ...interface{}) {
	if Level > NOT_EVEN_STDERR {
		fmt.Fprint(os.Stderr, fmt.Sprintf(format, elements...), "\n")
	}
}

func StdErrl(lambda func() string) {
	if Level > NOT_EVEN_STDERR {
		fmt.Fprint(os.Stderr, lambda(), "\n")
	}
}

func IsDebugEnabled() bool {
	return Level == DEBUG
}

func Debug(a ...interface{}) {
	if Level == DEBUG {
		fmt.Fprint(os.Stdout, line(DEBUG, a...))
	}
}

func Debugf(format string, elements ...interface{}) {
	if Level == DEBUG {
		fmt.Fprint(os.Stdout, linef(DEBUG, format, elements...))
	}
}

func Debugl(lambda func() string) {
	if Level == DEBUG {
		fmt.Fprint(os.Stdout, line(DEBUG, lambda()))
	}
}

func Info(a ...interface{}) {
	if Level >= INFO {
		fmt.Fprint(os.Stdout, line(INFO, a...))
	}
}

func Infof(format string, elements ...interface{}) {
	if Level >= INFO {
		fmt.Fprint(os.Stdout, linef(INFO, format, elements...))
	}
}

func Infol(lambda func() string) {
	if Level >= INFO {
		fmt.Fprint(os.Stdout, line(INFO, lambda()))
	}
}

func Warn(a ...interface{}) {
	if Level >= WARN {
		fmt.Fprint(os.Stdout, line(WARN, a...))
	}
}

func Warnf(format string, elements ...interface{}) {
	if Level >= WARN {
		fmt.Fprint(os.Stdout, linef(WARN, format, elements...))
	}
}

func Warnl(lambda func() string) {
	if Level >= WARN {
		fmt.Fprint(os.Stdout, line(WARN, lambda()))
	}
}

func Error(a ...interface{}) {
	if Level >= ERROR {
		fmt.Fprint(os.Stdout, line(ERROR, a...))
	}
}

func Errorf(format string, elements ...interface{}) {
	if Level >= ERROR {
		fmt.Fprint(os.Stderr, linef(ERROR, format, elements...))
	}
}

func Errorl(lambda func() string) {
	if Level >= ERROR {
		fmt.Fprint(os.Stderr, line(ERROR, lambda()))
	}
}

func Fatal(a ...interface{}) {
	if Level >= FATAL {
		fmt.Fprint(os.Stderr, line(FATAL, a...))
	}
	ForFatal()
}

func Fatalf(format string, elements ...interface{}) {
	if Level >= FATAL {
		fmt.Fprint(os.Stderr, linef(FATAL, format, elements...))
	}
	ForFatal()
}

func Fatall(lambda func() string) {
	if Level >= FATAL {
		fmt.Fprint(os.Stderr, line(FATAL, lambda()))
	}
	ForFatal()
}
