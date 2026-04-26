package logger

import (
	"log"
	"os"
)

type Logger struct {
	*os.File
}

func NewLogger(filepath string) *Logger {
	var err error
	l := &Logger{}
	l.File, err = os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
	if err != nil {
		return nil
	}

	// Set the standard log package to write to this file instead of the screen
	log.SetOutput(l.File)

	// Optional: customize flags (date, time, file name)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	return l
}
