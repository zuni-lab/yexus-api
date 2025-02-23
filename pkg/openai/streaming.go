package openai

import (
	"context"
	"fmt"
	"net/http"

	"github.com/openai/openai-go"
	"github.com/zuni-lab/dexon-service/config"
)

const (
	sseEventFormat     = "event: %s\ndata: %s\n\n"
	threadEventName    = "thread"
	messageEventName   = "message"
	errorEventName     = "error"
	doneEventName      = "done"
	threadDataTemplate = "{\"thread_id\": \"%s\"}"
	errorDataTemplate  = "{\"error\": \"%s\"}"
	doneDataTemplate   = "{\"status\": \"completed\"}"
)

func Streaming(ctx context.Context, threadID string, message string, w http.ResponseWriter) error {
	flusher, ok := w.(http.Flusher)
	if !ok {
		return fmt.Errorf("streaming unsupported")
	}

	// Create thread if not provided
	if threadID == "" {
		var err error
		threadID, err = CreateThread(ctx)
		if err != nil {
			return err
		}
	}

	// Send message to OpenAI
	_, err := openaiClient.Beta.Threads.Messages.New(ctx, threadID, openai.BetaThreadMessageNewParams{
		Role: openai.F(openai.BetaThreadMessageNewParamsRoleUser), // Changed to User role
		Content: openai.F([]openai.MessageContentPartParamUnion{
			openai.TextContentBlockParam{
				Type: openai.F(openai.TextContentBlockParamTypeText),
				Text: openai.String(message),
			},
		}),
	})
	if err != nil {
		return err
	}

	stream := openaiClient.Beta.Threads.Runs.NewStreaming(ctx, threadID, openai.BetaThreadRunNewParams{
		AssistantID: openai.String(config.Env.OpenaiAssistantId),
	})
	defer stream.Close()

	// Send initial SSE message with threadId
	fmt.Fprintf(w, sseEventFormat, threadEventName, fmt.Sprintf(threadDataTemplate, threadID))
	flusher.Flush()

	for stream.Next() {
		evt := stream.Current()
		if delta, ok := evt.Data.(openai.MessageDeltaEvent); ok {
			if delta.Delta.Content != nil {
				chunk := delta.Delta.Content[0].Text.Value
				// Format as SSE message
				fmt.Fprintf(w, sseEventFormat, messageEventName, chunk)
				flusher.Flush()
			}
		}
	}

	if err := stream.Err(); err != nil {
		// Send error event
		fmt.Fprintf(w, sseEventFormat, errorEventName, fmt.Sprintf(errorDataTemplate, err.Error()))
		flusher.Flush()
		return err
	}

	// Send done event
	fmt.Fprintf(w, sseEventFormat, doneEventName, doneDataTemplate)
	flusher.Flush()

	return nil
}
