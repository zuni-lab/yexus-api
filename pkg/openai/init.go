package openai

import (
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/zuni-lab/yexus-api/config"
)

var (
	openaiClient *openai.Client
)

func Init() {
	openaiClient = openai.NewClient(
		option.WithAPIKey(config.Env.OpenaiApiKey),
	)
}
