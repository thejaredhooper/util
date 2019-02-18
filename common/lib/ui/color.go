package ui

import (
	"flag"
	"github.com/logrusorgru/aurora"
)

// colorizer
var (
	COLOR = aurora.NewAurora(true)
)

func Flags() (cPtr *bool) {
	cPtr = flag.Bool("disableColors", false, "enable or disable colors")
	return
}

func CustomColors(disable bool) {
	COLOR = aurora.NewAurora(!disable)
}
