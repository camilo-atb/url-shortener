package application

import (
	"log"
	"time"
	"url_shortener/internal/domain"
)

type createShortLink struct { // ES un Input Port: Porque alguien desde afuera dice:
	repo          domain.ShortLinkRepository // Guardar links
	codeGenerator domain.CodeGenerator       // Generar códigos
	clock         domain.Clock               // Saber la hora
	idGen         domain.IDGenerator         // Generar ID
}

type CreateShortLinkI interface {
	Execute(url string, expiration *time.Time) (*domain.ShortLink, error)
}

func NewCreateShortLink(
	repo domain.ShortLinkRepository,
	codeGen domain.CodeGenerator,
	clock domain.Clock,
	idGen domain.IDGenerator,
) *createShortLink {
	return &createShortLink{
		repo:          repo,
		codeGenerator: codeGen,
		clock:         clock,
		idGen:         idGen,
	}
}

func (uc *createShortLink) Execute(
	url string,
	expiration *time.Time,
) (*domain.ShortLink, error) {
	code, err := uc.codeGenerator.Generate()

	log.Println("creating shortlink")
	if err != nil {
		return nil, err
	}

	link := &domain.ShortLink{
		ID:          uc.idGen.NewID(),
		OriginalURL: url,
		Code:        code,
		CreatedAt:   uc.clock.Now(),
		ExpiresAt:   expiration,
	}

	err = uc.repo.Save(link)

	if err != nil {
		return nil, err
	}
	log.Println("creating shortlink")
	return link, nil
}
