package ports

type EmbeddingsGeneratorInput struct {
	Text string
}

type EmbeddingsGeneratorOutput struct {
	Values []float32 `json:"embedding"`
}

type EmbeddingsGenerator interface {
	Generate(input EmbeddingsGeneratorInput) (*EmbeddingsGeneratorOutput, error)
}
