package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/schema"
	"github.com/subosito/gotenv"
)

func main() {
	ctx := context.Background()

	// 尝试加载 .env 文件，如果不存在则继续（因为可能在环境变量中设置了）
	if err := gotenv.Load(".env"); err != nil {
		fmt.Printf("Warning: could not load .env file: %v\n", err)
	}

	// 从环境变量获取配置值（gotenv.Load 会将 .env 文件中的值加载到环境变量中）
	apiKey := os.Getenv("API_KEY")
	model := os.Getenv("MODEL")
	baseURL := os.Getenv("BASE_URL")

	// 检查必需的环境变量
	if apiKey == "" {
		fmt.Println("Error: API_KEY environment variable is required")
		os.Exit(1)
	}
	if model == "" {
		fmt.Println("Error: MODEL environment variable is required")
		os.Exit(1)
	}
	if baseURL == "" {
		fmt.Println("Error: BASE_URL environment variable is required")
		os.Exit(1)
	}

	// Create a chat model instance (using OpenAI as an example)
	chatModel, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		APIKey:  apiKey,
		Model:   model,
		BaseURL: baseURL,
	})
	if err != nil {
		fmt.Printf("NewChatModel failed, err=%v\n", err)
		return
	}

	// 构建消息列表：先设置系统角色，然后添加用户消息
	messages := []*schema.Message{
		{
			Role:    "system",
			Content: "你是小哲AI助手",
		},
		{
			Role:    "user",
			Content: "今天星期几？天气怎么样",
		},
	}

	resp, err := chatModel.Generate(ctx, messages)
	if err != nil {
		fmt.Printf("Generate failed, err=%v\n", err)
		return
	}

	// 输出响应内容
	if resp != nil {
		content, err:= resp.Format(ctx, nil, schema.FString)
		if err != nil {
			fmt.Printf("Format failed, err=%v\n", err)
			return
		}
		fmt.Printf("AI Response: %v\n", content[0].Content)
	}

}
