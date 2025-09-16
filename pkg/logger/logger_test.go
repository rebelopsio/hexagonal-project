package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"
	"testing"
)

func TestDebugLogging(t *testing.T) {
	tests := []struct {
		name    string
		level   Level
		message string
		args    []any
		wantLog bool
		wantMsg string
	}{
		{
			name:    "debug message logged when level is debug",
			level:   LevelDebug,
			message: "test debug message",
			args:    []any{"key", "value"},
			wantLog: true,
			wantMsg: "test debug message",
		},
		{
			name:    "debug message not logged when level is info",
			level:   LevelInfo,
			message: "test debug message",
			args:    []any{"key", "value"},
			wantLog: false,
		},
		{
			name:    "debug message with no args",
			level:   LevelDebug,
			message: "simple debug",
			args:    nil,
			wantLog: true,
			wantMsg: "simple debug",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			logger := New(&buf, tt.level, "test-service", nil)

			logger.Debug(context.Background(), tt.message, tt.args...)

			output := buf.String()

			if !tt.wantLog {
				if output != "" {
					t.Errorf("expected no log output, but got: %s", output)
				}
				return
			}

			// Should have output when wantLog is true
			if output == "" {
				t.Error("expected log output, but got empty string")
				return
			}

			// Parse the JSON log entry
			var logEntry map[string]any
			if err := json.Unmarshal([]byte(strings.TrimSpace(output)), &logEntry); err != nil {
				t.Fatalf("failed to parse log JSON: %v", err)
			}

			// Verify log structure
			if msg, ok := logEntry["msg"].(string); !ok || msg != tt.wantMsg {
				t.Errorf("expected message %q, got %q", tt.wantMsg, msg)
			}

			if level, ok := logEntry["level"].(string); !ok || level != "DEBUG" {
				t.Errorf("expected level DEBUG, got %q", level)
			}

			if service, ok := logEntry["service"].(string); !ok || service != "test-service" {
				t.Errorf("expected service test-service, got %q", service)
			}

			// Verify timestamp exists
			if _, ok := logEntry["time"]; !ok {
				t.Error("expected time field in log entry")
			}

			// Verify file source exists
			if _, ok := logEntry["file"]; !ok {
				t.Error("expected file field in log entry")
			}

			// Verify args were logged if provided
			if len(tt.args) > 0 && tt.args[0] == "key" {
				if value, ok := logEntry["key"].(string); !ok || value != "value" {
					t.Errorf("expected key=value in log entry, got %v", value)
				}
			}
		})
	}
}

func TestDebugLoggingWithTraceID(t *testing.T) {
	var buf bytes.Buffer

	// Mock trace ID function
	traceIDFn := func(ctx context.Context) string {
		return "test-trace-123"
	}

	logger := New(&buf, LevelDebug, "test-service", traceIDFn)
	logger.Debug(context.Background(), "test message")

	output := buf.String()
	if output == "" {
		t.Fatal("expected log output")
	}

	var logEntry map[string]any
	if err := json.Unmarshal([]byte(strings.TrimSpace(output)), &logEntry); err != nil {
		t.Fatalf("failed to parse log JSON: %v", err)
	}

	if traceID, ok := logEntry["trace_id"].(string); !ok || traceID != "test-trace-123" {
		t.Errorf("expected trace_id=test-trace-123, got %q", traceID)
	}
}

func TestInfoLogging(t *testing.T) {
	tests := []struct {
		name    string
		level   Level
		message string
		args    []any
		wantLog bool
		wantMsg string
	}{
		{
			name:    "debug message logged when level is debug",
			level:   LevelDebug,
			message: "test debug message",
			args:    []any{"key", "value"},
			wantLog: true,
			wantMsg: "test debug message",
		},
		{
			name:    "debug message not logged when level is info",
			level:   LevelInfo,
			message: "test debug message",
			args:    []any{"key", "value"},
			wantLog: false,
		},
		{
			name:    "debug message with no args",
			level:   LevelDebug,
			message: "simple debug",
			args:    nil,
			wantLog: true,
			wantMsg: "simple debug",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			logger := New(&buf, tt.level, "test-service", nil)

			logger.Debug(context.Background(), tt.message, tt.args...)

			output := buf.String()

			if !tt.wantLog {
				if output != "" {
					t.Errorf("expected no log output, but got: %s", output)
				}
				return
			}

			// Should have output when wantLog is true
			if output == "" {
				t.Error("expected log output, but got empty string")
				return
			}

			// Parse the JSON log entry
			var logEntry map[string]any
			if err := json.Unmarshal([]byte(strings.TrimSpace(output)), &logEntry); err != nil {
				t.Fatalf("failed to parse log JSON: %v", err)
			}

			// Verify log structure
			if msg, ok := logEntry["msg"].(string); !ok || msg != tt.wantMsg {
				t.Errorf("expected message %q, got %q", tt.wantMsg, msg)
			}

			if level, ok := logEntry["level"].(string); !ok || level != "DEBUG" {
				t.Errorf("expected level DEBUG, got %q", level)
			}

			if service, ok := logEntry["service"].(string); !ok || service != "test-service" {
				t.Errorf("expected service test-service, got %q", service)
			}

			// Verify timestamp exists
			if _, ok := logEntry["time"]; !ok {
				t.Error("expected time field in log entry")
			}

			// Verify file source exists
			if _, ok := logEntry["file"]; !ok {
				t.Error("expected file field in log entry")
			}

			// Verify args were logged if provided
			if len(tt.args) > 0 && tt.args[0] == "key" {
				if value, ok := logEntry["key"].(string); !ok || value != "value" {
					t.Errorf("expected key=value in log entry, got %v", value)
				}
			}
		})
	}
}
