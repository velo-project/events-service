package ai

import (
	"context"
	"time"

	"github.com/google/generative-ai-go/genai"
)

type EmbeddingsGeneratorInput struct {
	Text string
}

type EmbeddingsGeneratorOutput struct {
	Values []float32 `json:"embedding"`
}

type EmbeddingsGenerator interface {
	Generate(input EmbeddingsGeneratorInput) (*EmbeddingsGeneratorOutput, error)
}

type embeddingsGenerator struct {
	Model *genai.EmbeddingModel
}

func NewEmbeddingsGeneratorOutput(model *genai.EmbeddingModel) EmbeddingsGenerator {
	return &embeddingsGenerator{
		Model: model,
	}
}

func (e embeddingsGenerator) Generate(input EmbeddingsGeneratorInput) (*EmbeddingsGeneratorOutput, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	res, err := e.Model.EmbedContent(ctx, genai.Text(input.Text))
	if err != nil {
		return nil, err
	}

	embeddings := &EmbeddingsGeneratorOutput{
		Values: res.Embedding.Values,
	}

	return embeddings, nil
}
