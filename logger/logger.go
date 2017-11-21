package logger

import "os"
import "log"

var info, errs, fatal *log.Logger

func init() {
	info  = log.New(os.Stdout, "[INFO ] ", log.LstdFlags)
	errs  = log.New(os.Stderr, "[ERROR] ", log.LstdFlags)
	fatal = log.New(os.Stderr, "[FATAL] ", log.LstdFlags)
}

func Info(v ...interface{}) {
	print(info, v)
}

func Error(v ...interface{}) {
	print(errs, v)
}

func Fatal(v ...interface{}) {
	print(fatal, v)
	os.Exit(1)
}

func print(log *log.Logger, v []interface{}) {
	if w, ok := v[0].(string); ok {
		log.Printf(w, v[1:]...)
	} else {
		log.Println(v...)
	}
}