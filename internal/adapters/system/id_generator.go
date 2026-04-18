package system

import "github.com/google/uuid"

type idGenerator struct{}

func NewIDGenerator() *idGenerator {
	return &idGenerator{}
}

func (g *idGenerator) NewID() string {
	return uuid.NewString()
}
