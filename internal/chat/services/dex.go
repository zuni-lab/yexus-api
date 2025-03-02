package services

import (
	"context"
	"encoding/json"
	"fmt"
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

	cryptoData := []map[string]string{
		{
			"token_name": "WBTC",
			"price":      "90000",
		},
		{
			"token_name": "WETH",
			"price":      "3000",
		},
		{
			"token_name": "WSOL",
			"price":      "200",
		},
	}

	cryptoDataString, errJson := json.Marshal(cryptoData)
	if errJson != nil {
		return fmt.Errorf("failed to marshal crypto data: %w", errJson)
	}

	// Create a new thread if threadId is not provided
	if input.ThreadID == "" {
		threadID, _ = openai.CreateThreadWithData(ctx, string(cryptoDataString))
	} else {
		threadID = input.ThreadID
	}

	log.Debug().Msgf("threadID: %s", threadID)

	return openai.Streaming(ctx, threadID, input.Message, w)
}

type GetMessagesListParams struct {
	ThreadID string `json:"thread_id"`
}

func GetMessagesList(ctx context.Context, input GetMessagesListParams, w http.ResponseWriter) error {
	messages, err := openai.GetMessagesList(ctx, input.ThreadID)
	if err != nil {
		return fmt.Errorf("failed to get messages list: %w", err)
	}

	return json.NewEncoder(w).Encode(messages)
}
