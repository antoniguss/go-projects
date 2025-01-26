package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func SetupBasicRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logger, ok := r.Context().Value("logger").(*slog.Logger)
		if !ok || logger == nil {
			logger = slog.Default() // Use a default logger
			logger.Warn("logger not set in context, using default logger")
		}

		logger.Info("processing base path request")
	})

	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		logger, ok := r.Context().Value("logger").(*slog.Logger)
		if !ok || logger == nil {
			logger = slog.Default() // Use a default logger
			logger.Warn("logger not set in context, using default logger")
		}

		logger.Info("processing ping request")

		pongMsg, _ := json.Marshal(map[string]string{"message": "pong"})

		fmt.Fprintf(w, "%s", pongMsg)
	})

	mux.HandleFunc("POST /add", func(w http.ResponseWriter, r *http.Request) {
		logger, ok := r.Context().Value("logger").(*slog.Logger)
		if !ok || logger == nil {
			logger = slog.Default() // Use a default logger
			logger.Warn("logger not set in context, using default logger")
		}

		logger.Info("processing add request")

		data := make([]byte, r.ContentLength)
		r.Body.Read(data)

		input := &InputStruct{}
		err := json.Unmarshal(data, input)
		if err != nil {

			logger.Error("failed to process request",
				"error", err,
				"operation", "add",
				"input", input,
			)
			http.Error(w, fmt.Sprintf("invalid input (%v)", err), http.StatusBadRequest)
			return

		}

		if input.Number1 == nil || input.Number2 == nil {
			logger.Error("failed to process request",
				"error", "missing value",
				"operation", "add",
				"input", input,
			)
			http.Error(w, "invalid input (missing value)", http.StatusBadRequest)
			return
		}
		result := ResultStruct{Result: *input.Number1 + *input.Number2}
		resultMsg, _ := json.Marshal(result)
		fmt.Fprintf(w, "%s", resultMsg)
	})

	mux.HandleFunc("POST /subtract", func(w http.ResponseWriter, r *http.Request) {
		logger, ok := r.Context().Value("logger").(*slog.Logger)
		if !ok || logger == nil {
			logger = slog.Default() // Use a default logger
			logger.Warn("logger not set in context, using default logger")
		}

		logger.Info("processing subtract request")

		data := make([]byte, r.ContentLength)
		r.Body.Read(data)

		input := &InputStruct{}
		err := json.Unmarshal(data, input)
		if err != nil {

			logger.Error("failed to process request",
				"error", err,
				"operation", "subtract",
				"input", input,
			)
			http.Error(w, fmt.Sprintf("invalid input (%v)", err), http.StatusBadRequest)
			return

		}

		if input.Number1 == nil || input.Number2 == nil {
			logger.Error("failed to process request",
				"error", "missing value",
				"operation", "subtract",
				"input", input,
			)
			http.Error(w, "invalid input (missing value)", http.StatusBadRequest)
			return
		}
		result := ResultStruct{Result: *input.Number1 - *input.Number2}
		resultMsg, _ := json.Marshal(result)
		fmt.Fprintf(w, "%s", resultMsg)
	})

	mux.HandleFunc("POST /multiply", func(w http.ResponseWriter, r *http.Request) {
		logger, ok := r.Context().Value("logger").(*slog.Logger)
		if !ok || logger == nil {
			logger = slog.Default() // Use a default logger
			logger.Warn("logger not set in context, using default logger")
		}

		logger.Info("processing multiply request")

		data := make([]byte, r.ContentLength)
		r.Body.Read(data)

		input := &InputStruct{}
		err := json.Unmarshal(data, input)
		if err != nil {

			logger.Error("failed to process request",
				"error", err,
				"operation", "multiply",
				"input", input,
			)
			http.Error(w, fmt.Sprintf("invalid input (%v)", err), http.StatusBadRequest)
			return

		}

		if input.Number1 == nil || input.Number2 == nil {
			logger.Error("failed to process request",
				"error", "missing value",
				"operation", "multiply",
				"input", input,
			)
			http.Error(w, "invalid input (missing value)", http.StatusBadRequest)
			return
		}
		result := ResultStruct{Result: *input.Number1 * *input.Number2}
		resultMsg, _ := json.Marshal(result)
		fmt.Fprintf(w, "%s", resultMsg)
	})

	mux.HandleFunc("POST /divide", func(w http.ResponseWriter, r *http.Request) {
		logger, ok := r.Context().Value("logger").(*slog.Logger)
		if !ok || logger == nil {
			logger = slog.Default() // Use a default logger
			logger.Warn("logger not set in context, using default logger")
		}

		logger.Info("processing divide request")
		data := make([]byte, r.ContentLength)
		r.Body.Read(data)

		input := &InputStruct{}
		err := json.Unmarshal(data, input)
		if err != nil {

			logger.Error("failed to process request",
				"error", err,
				"operation", "divide",
				"input", input,
			)
			http.Error(w, fmt.Sprintf("invalid input (%v)", err), http.StatusBadRequest)
			return

		}

		if input.Number1 == nil || input.Number2 == nil {
			logger.Error("failed to process request",
				"error", "missing value",
				"operation", "divide",
				"input", input,
			)
			http.Error(w, "invalid input (missing value)", http.StatusBadRequest)
			return
		}

		if *input.Number2 == 0 {
			logger.Error("failed to process request",
				"error", "division by zero",
				"operation", "divide",
				"input", input,
			)
			http.Error(w, "invalid input (division by zero)", http.StatusBadRequest)
			return
		}

		result := ResultStruct{Result: *input.Number1 / *input.Number2}
		resultMsg, _ := json.Marshal(result)
		fmt.Fprintf(w, "%s", resultMsg)
	})

	return mux
}

func SetupLoggingMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqLogger := slog.With(
			"method", r.Method,
			"path", r.URL.Path,
			"remote_addr", r.RemoteAddr,
		)

		ctx := context.WithValue(r.Context(), "logger", reqLogger)
		r = r.WithContext(ctx)

		reqLogger.Info("request started")

		start := time.Now()
		next.ServeHTTP(w, r)

		reqLogger.Info("request completed", "duration_ms", time.Since(start).Milliseconds())
	})
}
