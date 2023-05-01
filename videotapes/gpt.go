package videotapes

import (
	"context"
	"errors"

	"fmt"
	openai "github.com/sashabaranov/go-openai"

)

type GptConnection struct {
	ApiToken string
}


func (conn *GptConnection) GetResponse(prompt string) (string, error) {
	client := openai.NewClient(conn.ApiToken)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	
	

	if err != nil {
		fmt.Println(err)
		return "Error" , errors.New("Error")
	}

	return resp.Choices[0].Message.Content, nil
}
