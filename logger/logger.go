package logger

import (
	"log"
	"net/http"
	"os"
)

type HandlerLogger struct {
	handler http.Handler
}

type Logger struct {
	message string
	params  interface{}
}

type ErrLogger struct {
	message string
	error   error
}

func (l *HandlerLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	file, err := os.OpenFile("http.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Println("%s %s %v", r.Method, r.URL.Path)
}

func HttpLogger(handlerToWrap http.Handler) *HandlerLogger {
	return &HandlerLogger{handlerToWrap}
}

func (l *Logger) Log() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Printf(l.message, l.params)
}

func CommonLogger(message string, params interface{}) *Logger {
	return &Logger{message, params}
}

func (l *ErrLogger) Log() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Printf(l.message, l.error)
}

func ErrorLogger(message string, error error) *ErrLogger {
	return &ErrLogger{message, error}
}
