package logger

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	*os.File
}

func NewLogger(logToFile bool) *Logger {
	l := &Logger{}
	var err error

	if logToFile {
		l.File, err = os.CreateTemp("/tmp", "dbactivity")
		if err != nil {
			log.Fatal("Failed to create log file:", err)
			return nil
		}
		log.SetOutput(l.File)
	} else {
		log.SetOutput(io.Discard)
	}

	// Customize flags (date, time, file name)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)

	return l
}
