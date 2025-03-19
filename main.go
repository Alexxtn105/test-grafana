package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	addr := ":8080"

	// Configure slog to use JSON format and output to stderr
	logger := slog.New(slog.NewJSONHandler(os.Stderr,
		&slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	//models.ConnectDatabase()
	//models.DBMigrate()

	mux := http.NewServeMux()
	mux.HandleFunc("/", firstMiddleware(secondMiddleware(HelloHandler)))
	//mux.HandleFunc("/blogs", firstMiddleware(secondMiddleware(controllers.BlogsIndex)))

	muxWithLogger := NewLogger(mux)

	slog.Info("Server is starting", slog.String("address", addr))
	http.ListenAndServe(addr, muxWithLogger)

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func firstMiddleware(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("In Middleware - Before handler")
		f(w, r) // original function call
		slog.Info("In Middleware - After handler")
	}
}

func secondMiddleware(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("In Middleware 2 - Before handler")
		f(w, r) // original function call
		slog.Info("In Middleware 2 - After handler")
	}
}

// Logger Middleware handler to log requests
type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	slog.Info("Request completed",
		slog.String("method", r.Method),
		slog.String("path", r.URL.Path),
		slog.Duration("duration", time.Since(start)),
	)
}

// NewLogger constructs a new Logger middleware handler
func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}
