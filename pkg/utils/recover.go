package utils

import (
	"fmt"
	"runtime/debug"

	"github.com/rs/zerolog/log"
)

func Recover(module string, event interface{}, message string) {
	if r := recover(); r != nil {
		err := fmt.Errorf("panic recovered in %s: %v\n%s", module, r, debug.Stack())
		log.Error().
			Err(err).
			Interface("event", event).
			Msg(message)
	}
}
