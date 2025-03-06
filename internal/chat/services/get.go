package services

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/openai/openai-go"
	"github.com/zuni-lab/dexon-service/pkg/db"
	openAI "github.com/zuni-lab/dexon-service/pkg/openai"
	"github.com/zuni-lab/dexon-service/pkg/utils"
)

type GetThreadDetailsParams struct {
	ThreadID    string `json:"thread_id"`
	UserAddress string `json:"user_address" validate:"eth_addr"`
}

type Message struct {
	Role      openai.MessageRole `json:"role"`
	Content   string             `json:"text"`
	CreatedAt int64              `json:"created_at"`
}

type GetThreadDetailsResponse struct {
	Message    []Message `json:"message"`
	ThreadID   string    `json:"thread_id"`
	ThreadName string    `json:"thread_name"`
	UpdatedAt  int64     `json:"updated_at"`
}

func GetThreadDetails(ctx context.Context, input GetThreadDetailsParams) (*GetThreadDetailsResponse, error) {
	thread, err := db.DB.GetChatThread(ctx, db.GetChatThreadParams{
		ThreadID:    input.ThreadID,
		UserAddress: utils.NormalizeAddress(input.UserAddress),
	})
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, "thread not found")
	}

	messages, err := openAI.GetMessagesList(ctx, input.ThreadID)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// not get the last message
	if len(messages) > 0 {
		messages = messages[:len(messages)-1]
	}

	// process messages
	filtered := processMessage(messages)

	// convert updated_at to unix timestamp
	updatedAt := thread.UpdatedAt.Time.Unix()

	return &GetThreadDetailsResponse{
		Message:    filtered,
		ThreadID:   input.ThreadID,
		ThreadName: thread.ThreadName,
		UpdatedAt:  updatedAt,
	}, nil
}

type GetThreadListParams struct {
	UserAddress string `json:"user_address" validate:"eth_addr"`
	Limit       int32  `json:"limit" validate:"gte=0"`
	Offset      int32  `json:"offset" validate:"gte=0"`
}

type ChatThread struct {
	ID          int64  `json:"id"`
	ThreadID    string `json:"thread_id"`
	UserAddress string `json:"user_address"`
	ThreadName  string `json:"thread_name"`
	UpdatedAt   int64  `json:"updated_at"`
}

type GetThreadListResponse struct {
	Threads []ChatThread `json:"threads"`
	Count   int64        `json:"count"`
}

func GetThreadList(ctx context.Context, input GetThreadListParams) (*GetThreadListResponse, error) {
	result, err := db.DB.ListThreadsTx(ctx, db.GetChatThreadsParams{
		Limit:       input.Limit,
		Offset:      input.Offset,
		UserAddress: utils.NormalizeAddress(input.UserAddress),
	})
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	newThreads := utils.Map(result.Threads, func(thread db.ChatThread) ChatThread {
		return ChatThread{
			ID:          thread.ID,
			ThreadID:    thread.ThreadID,
			UserAddress: thread.UserAddress,
			ThreadName:  thread.ThreadName,
			UpdatedAt:   thread.UpdatedAt.Time.Unix(),
		}
	})

	return &GetThreadListResponse{
		Threads: newThreads,
		Count:   result.Count,
	}, nil
}

func processMessage(messages []openai.Message) []Message {
	var result []Message
	for _, message := range messages {
		content := ""
		if len(message.Content) > 0 {
			content = message.Content[0].Text.Value
		}
		result = append(result, Message{
			Role:      message.Role,
			Content:   content,
			CreatedAt: message.CreatedAt,
		})
	}
	return result
}
