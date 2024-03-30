package masked_test

import (
	"log/slog"
	"strings"
	"testing"

	"github.com/pedrolemes/sensitive/internal/masked"
)

type testLogWriter struct {
	b []byte
}

func (w *testLogWriter) Clear() {
	w.b = nil
}

func (w *testLogWriter) Write(p []byte) (n int, err error) {
	w.b = p
	return len(p), nil
}

func TestLogValue(t *testing.T) {
	t.Run("ensure slog logs the masked value", func(t *testing.T) {
		writer := &testLogWriter{}

		loggers := []*slog.Logger{
			slog.New(slog.NewTextHandler(writer, nil)),
			slog.New(slog.NewJSONHandler(writer, nil)),
		}

		for _, logger := range loggers {

			t.Run("logs directly", func(t *testing.T) {
				writer.Clear()
				m := masked.NewMaskedText(sensitive, testMaskFunc)

				logger.Info("log", "value", m)
				got := string(writer.b)

				if strings.Contains(got, sensitive) {
					t.Errorf("expected %q to not contain %q", got, sensitive)
				}

				if !strings.Contains(got, handled) {
					t.Errorf("expected %q to contain %q", got, handled)
				}
			})

			t.Run("logs as a field", func(t *testing.T) {
				writer.Clear()

				type someStruct struct {
					Value masked.MaskedText
				}

				some := someStruct{Value: masked.NewMaskedText(sensitive, testMaskFunc)}

				logger.Info("log", "value", some)
				got := string(writer.b)

				if strings.Contains(got, sensitive) {
					t.Errorf("expected %q to not contain %q", got, sensitive)
				}

				if !strings.Contains(got, handled) {
					t.Errorf("expected %q to contain %q", got, handled)
				}
			})
		}
	})
}
