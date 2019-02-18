package lib

import (
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"strings"
	"time"
)

var (
	DEBUG = false
	DB    = "health-profile"

	LineFormat = "\n %-25s | %30s | %30s "
	C1         = 25
	C2         = 30
	C3         = 30

	LineBreak = "\n" + strings.Repeat("-", len(fmt.Sprintf(LineFormat, "", "", "")))
	Success   = `✓`
	Failure   = `✗`
)

// panics!
func Fatal(message string, err error) {
	Printf("%s: %s", message, err.Error())
}

func CaseInsensitive(value string) bson.M {
	return bson.M{
		"$regex":   value,
		"$options": "i",
	}
}

func Time(message string, f func()) {
	executionTime := time.Now()
	f()
	fmt.Printf("\t(%s took %v)", message, Elapsed(executionTime))
}

func Elapsed(start time.Time) string {
	e := time.Since(start)
	return fmt.Sprintf("%s", e)
}

func Truncate(str string, num int) string {
	truncated := str
	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		truncated = str[0:num] + "..."
	}
	return truncated
}

func Debug(flag bool) {
	DEBUG = flag
}

func Silent() bool {
	return !DEBUG
}

func Printf(format string, a ...interface{}) (n int, err error) {
	if ! Silent() {
		n, err = fmt.Printf(format, a...)
	}

	return
}

func Println(a ...interface{}) (n int, err error) {
	if ! Silent() {
		n, err = fmt.Println(a...)
	}

	return
}

func Print(a ...interface{}) (n int, err error) {
	if ! Silent() {
		n, err = fmt.Print(a...)
	}

	return
}
