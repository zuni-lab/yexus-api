package services

import (
	"context"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/zuni-lab/dexon-service/pkg/openai"
)

type ChatParams struct {
	Message  string `json:"message"`
	ThreadID string `json:"thread_id,omitempty"`
}

type ChatResponse struct {
	Message  string `json:"message"`
	ThreadID string `json:"thread_id"`
}

func ChatDex(ctx context.Context, input ChatParams, w http.ResponseWriter) error {
	var threadID string

	// Create a new thread if threadId is not provided
	if input.ThreadID == "" {
		threadID, _ = openai.CreateThread(ctx)
	} else {
		threadID = input.ThreadID
	}

	log.Info().Msgf("threadID: %s", threadID)

	return openai.Streaming(ctx, threadID, input.Message, w)
}
