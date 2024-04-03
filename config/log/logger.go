package logger

import (
	"fmt"
	"log"
	"os"
)

func Logger_init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func Info(format string, a ...any) {
	log.SetPrefix("[INFO] ")
	log.Output(2, fmt.Sprintf(format, a...))
}
func Debug(format string, a ...any) {
	log.SetPrefix("[DEBUG] ")
	log.Output(2, fmt.Sprintf(format, a...))
}
func Warn(format string, a ...any) {
	log.SetPrefix("[WARN] ")
	log.Output(2, fmt.Sprintf(format, a...))
}
func Error(err error) {
	log.SetPrefix("[ERROR] ")
	log.Output(2, err.Error())
	os.Exit(1)
}
func Fatal(err error) {
	log.SetPrefix("[FATAL] ")
	log.Output(2, err.Error())
	os.Exit(1)
}
func Assert(err error) {
	if err != nil {
		log.SetPrefix("[FATAL] ")
		log.Output(2, err.Error())
		os.Exit(1)
	}
}
