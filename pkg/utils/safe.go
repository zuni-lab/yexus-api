package utils

import (
	"context"
	"fmt"
	"runtime/debug"

	"github.com/rs/zerolog/log"
)

// SafeExecute runs a function with panic recovery
// Returns any error from the function or from a panic
func SafeExecute(ctx context.Context, fn func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			stack := debug.Stack()
			log.Error().
				Interface("panic", r).
				Str("stack", string(stack)).
				Msg("Recovered from panic in handler")

			// Convert panic to error
			switch x := r.(type) {
			case string:
				err = fmt.Errorf("panic: %s", x)
			case error:
				err = fmt.Errorf("panic: %w", x)
			default:
				err = fmt.Errorf("panic: %v", x)
			}
		}
	}()

	return fn()
}
