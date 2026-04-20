package system

import (
	"url_shortener/internal/domain"

	"github.com/google/uuid"
)

type idGenerator struct{}

func NewIDGenerator() domain.IDGenerator {
	return &idGenerator{}
}

func (g *idGenerator) NewID() string {
	return uuid.NewString()
}
