package main

import (
	"github.com/rs/zerolog/log"
)

// logger implements the paho.Logger interface
type logger struct {
	prefix string
}

// Println is the library provided NOOPLogger's
// implementation of the required interface function()
func (l logger) Println(v ...interface{}) {
	log.Debug().Str("service", l.prefix).Fields(v)
}

// Printf is the library provided NOOPLogger's
// implementation of the required interface function(){}
func (l logger) Printf(format string, v ...interface{}) {
	log.Debug().Str("service", l.prefix).Msgf(format, v...)
}
