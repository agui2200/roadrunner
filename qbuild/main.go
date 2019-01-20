package main

import (
	rr "github.com/agui2200/roadrunner/cmd/rr/cmd"
	"github.com/sirupsen/logrus"
	// -packages- //
	// -commands- //
)

func main() {
	// -register- //
	rr.Logger.Formatter = &logrus.TextFormatter{ForceColors: true}
	rr.Execute()
}
