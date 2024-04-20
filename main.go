package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func main() {
	modelName := flag.String("model", "", "ollama model name")
	query := flag.String("query", "", "query to the model")
	stream := flag.Bool("stream", false, "stream output")
	flag.Parse()

	llm, err := ollama.New(ollama.WithModel(*modelName))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	if !*stream {
		completion, err := llms.GenerateFromSinglePrompt(ctx, llm, *query)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Response:\n", completion)
	} else {
		_, err = llms.GenerateFromSinglePrompt(ctx, llm, *query,
			llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
				fmt.Printf("chunk len=%d: %s\n", len(chunk), chunk)
				return nil
			}))
		if err != nil {
			log.Fatal(err)
		}
	}

}
