package ai

import (
	"context"
	"time"

	"github.com/google/generative-ai-go/genai"
	"gitlab.com/velo-company/services/events-service/internal/core/ports"
)

type embeddingsGenerator struct {
	Model *genai.EmbeddingModel
}

func NewEmbeddingsGeneratorOutput(model *genai.EmbeddingModel) ports.EmbeddingsGenerator {
	return &embeddingsGenerator{
		Model: model,
	}
}

func (e embeddingsGenerator) Generate(input ports.EmbeddingsGeneratorInput) (*ports.EmbeddingsGeneratorOutput, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	res, err := e.Model.EmbedContent(ctx, genai.Text(input.Text))
	if err != nil {
		return nil, err
	}

	embeddings := &ports.EmbeddingsGeneratorOutput{
		Values: res.Embedding.Values,
	}

	return embeddings, nil
}
