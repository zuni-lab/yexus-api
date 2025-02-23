package openai

import (
	"context"

	"github.com/openai/openai-go"
)

func CreateThread(ctx context.Context) (string, error) {
	thread, err := openaiClient.Beta.Threads.New(ctx, openai.BetaThreadNewParams{})
	if err != nil {
		return "", err
	}
	return thread.ID, nil
}

// func CreateThreadWithData(ctx context.Context, data string) (string, error) {
// 	thread, err := openaiClient.Beta.Threads.New(ctx, openai.BetaThreadNewParams{})
// 	if err != nil {
// 		return "", err
// 	}
// 	_, err = openaiClient.Beta.Threads.Messages.New(ctx, thread.ID, openai.BetaThreadMessageNewParams{
// 		Role: openai.F(openai.BetaThreadMessageNewParamsRoleAssistant),
// 		Content: openai.F([]openai.MessageContentPartParamUnion{
// 			openai.TextContentBlockParam{
// 				Type: openai.F(openai.TextContentBlockParamTypeText),
// 				Text: openai.String(assistantMessage + data),
// 			},
// 		}),
// 	})
// 	if err != nil {
// 		return "", err
// 	}

// 	return thread.ID, nil
// }
