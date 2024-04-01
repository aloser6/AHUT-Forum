package logger

import (
	"log"
)

func Logger_init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func Info(format string, a ...any) {
	log.SetPrefix("[INFO] ")
	log.Printf(format, a...)
}
func Debug(format string, a ...any) {
	log.SetPrefix("[DEBUG] ")
	log.Printf(format, a...)
}
func Warn(format string, a ...any) {
	log.SetPrefix("[WARN] ")
	log.Printf(format, a...)
}
func Error(err error) {
	log.SetPrefix("[ERROR] ")
	log.Fatal(err)
}
func Fatal(err error) {
	log.SetPrefix("[FATAL] ")
	log.Fatal(err)
}
func Assert(err error) {
	if err != nil {
		log.SetPrefix("[FATAL] ")
		log.Fatal(err)
	}
}
