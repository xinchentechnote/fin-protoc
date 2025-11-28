package parser

import (
	"github.com/xinchentechnote/fin-protoc/internal/model"
)

// Generator interface for extensibility
type Generator interface {
	Generate(binModel *model.BinaryModel) (map[string][]byte, error)
	GetConfig() *model.Configuration
}
